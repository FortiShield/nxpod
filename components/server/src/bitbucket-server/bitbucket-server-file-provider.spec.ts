/**
 * Copyright (c) 2022 Nxpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { Repository, User } from "@nxpod/nxpod-protocol";
import { ifEnvVarNotSet } from "@nxpod/nxpod-protocol/lib/util/skip-if";
import { Container, ContainerModule } from "inversify";
import { retries, skip, suite, test, timeout } from "@testdeck/mocha";
import { expect } from "chai";
import { NxpodHostUrl } from "@nxpod/nxpod-protocol/lib/util/nxpod-host-url";
import { BitbucketServerFileProvider } from "./bitbucket-server-file-provider";
import { AuthProviderParams } from "../auth/auth-provider";
import { BitbucketServerContextParser } from "./bitbucket-server-context-parser";
import { BitbucketServerTokenHelper } from "./bitbucket-server-token-handler";
import { TokenService } from "../user/token-service";
import { Config } from "../config";
import { TokenProvider } from "../user/token-provider";
import { BitbucketServerApi } from "./bitbucket-server-api";
import { HostContextProvider } from "../auth/host-context-provider";

@suite(timeout(10000), retries(1), skip(ifEnvVarNotSet("NXPOD_TEST_TOKEN_BITBUCKET_SERVER")))
class TestBitbucketServerFileProvider {
    protected service: BitbucketServerFileProvider;
    protected user: User;

    static readonly AUTH_HOST_CONFIG: Partial<AuthProviderParams> = {
        id: "MyBitbucketServer",
        type: "BitbucketServer",
        verified: true,
        description: "",
        icon: "",
        host: "bitbucket.nxpod-self-hosted.com",
        oauth: {
            callBackUrl: "",
            clientId: "not-used",
            clientSecret: "",
            tokenUrl: "",
            scope: "",
            authorizationUrl: "",
        },
    };

    public before() {
        const container = new Container();
        container.load(
            new ContainerModule((bind, unbind, isBound, rebind) => {
                bind(BitbucketServerFileProvider).toSelf().inSingletonScope();
                bind(BitbucketServerContextParser).toSelf().inSingletonScope();
                bind(AuthProviderParams).toConstantValue(TestBitbucketServerFileProvider.AUTH_HOST_CONFIG);
                bind(BitbucketServerTokenHelper).toSelf().inSingletonScope();
                // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
                bind(TokenService).toConstantValue({
                    createNxpodToken: async () => ({ token: { value: "foobar123-token" } }),
                } as any);
                bind(Config).toConstantValue({
                    hostUrl: new NxpodHostUrl("https://nxpod.khulnasoft.com"),
                });
                bind(TokenProvider).toConstantValue(<TokenProvider>{
                    getTokenForHost: async () => {
                        return {
                            value: process.env["NXPOD_TEST_TOKEN_BITBUCKET_SERVER"] || "undefined",
                            scopes: [],
                        };
                    },
                });
                bind(BitbucketServerApi).toSelf().inSingletonScope();
                bind(HostContextProvider).toConstantValue({
                    get: (hostname: string) => {
                        authProvider: {
                            ("BBS");
                        }
                    },
                });
            }),
        );
        this.service = container.get(BitbucketServerFileProvider);
        this.user = {
            creationDate: "",
            id: "user1",
            identities: [
                {
                    authId: "user1",
                    authName: "AlexTugarev",
                    authProviderId: "MyBitbucketServer",
                },
            ],
        };
    }

    @test async test_getNxpodFileContent_ok() {
        const result = await this.service.getNxpodFileContent(
            // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
            {
                revision: "master",
                repository: <Repository>{
                    cloneUrl: "https://bitbucket.nxpod-self-hosted.com/projects/FOO/repos/repo123",
                    webUrl: "https://bitbucket.nxpod-self-hosted.com/projects/FOO/repos/repo123",
                    name: "repo123",
                    repoKind: "projects",
                    owner: "FOO",
                },
            } as any,
            this.user,
        );
        expect(result).not.to.be.empty;
        expect(result).to.contain("tasks:");
    }

    @test async test_getLastChangeRevision_ok() {
        const result = await this.service.getLastChangeRevision(
            {
                owner: "FOO",
                name: "repo123",
                repoKind: "projects",
                revision: "foo",
                host: "bitbucket.nxpod-self-hosted.com",
                cloneUrl: "https://bitbucket.nxpod-self-hosted.com/projects/FOO/repos/repo123",
                webUrl: "https://bitbucket.nxpod-self-hosted.com/projects/FOO/repos/repo123",
            } as Repository,
            "foo",
            this.user,
            "folder/sub/test.txt",
        );
        expect(result).to.equal("1384b6842d73b8705feaf45f3e8aa41f00529042");
    }
}

module.exports = new TestBitbucketServerFileProvider();
