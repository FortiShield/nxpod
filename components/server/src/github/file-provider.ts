/**
 * Copyright (c) 2020 Nxpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { injectable, inject } from "inversify";

import { FileProvider, MaybeContent } from "../repohost/file-provider";
import { Commit, User, Repository } from "@nxpod/nxpod-protocol";
import { GitHubRestApi } from "./api";
import { log } from "@nxpod/nxpod-protocol/lib/util/logging";

@injectable()
export class GithubFileProvider implements FileProvider {
    @inject(GitHubRestApi) protected readonly githubApi: GitHubRestApi;

    public async getNxpodFileContent(commit: Commit, user: User): Promise<MaybeContent> {
        const yamlVersion1 = await Promise.all([
            this.getFileContent(commit, user, ".nxpod.yml"),
            this.getFileContent(commit, user, ".nxpod"),
        ]);
        return yamlVersion1.filter((f) => !!f)[0];
    }

    public async getLastChangeRevision(
        repository: Repository,
        revisionOrBranch: string,
        user: User,
        path: string,
    ): Promise<string> {
        const commits = (
            await this.githubApi.run(user, (gh) =>
                gh.repos.listCommits({
                    owner: repository.owner,
                    repo: repository.name,
                    sha: revisionOrBranch,
                    // per_page: 1, // we need just the last one right?
                    path,
                }),
            )
        ).data;

        const lastCommit = commits && commits[0];
        if (!lastCommit) {
            throw new Error(`File ${path} does not exist in repository ${repository.owner}/${repository.name}`);
        }

        return lastCommit.sha;
    }

    public async getFileContent(commit: Commit, user: User, path: string) {
        if (!commit.revision) {
            return undefined;
        }

        const params = {
            owner: commit.repository.owner,
            repo: commit.repository.name,
            path,
            ref: commit.revision,
            headers: {
                accept: "application/vnd.github.raw",
            },
        };

        try {
            const response = await this.githubApi.run(user, (api) => api.repos.getContent(params));
            if (response.status === 200) {
                if (typeof response.data === "string") {
                    return response.data;
                }
                log.warn("GithubFileProvider.getFileContent – unexpected response type.", {
                    request: params,
                    response: {
                        headers: {
                            "content-encoding": response.headers["content-encoding"],
                            "content-type": response.headers["content-type"],
                        },
                        type: typeof response.data,
                    },
                });
            }
            return undefined;
        } catch (err) {
            log.debug("Failed to get Github file content", err, {
                request: params,
            });
        }
    }
}
