// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Update JetBrains Platform Version

import path from "path";
import { parseArgs } from "util";
import { pathToProjectRoot } from "./lib/common";
import { jbUpdatePlatformVersion } from "./lib/jb-update-platform-version";

const targetInfo = [
    {
        id: 1,
        taskName: "Latest Backend Plugin",
        productCode: "IIU",
        productType: "eap,rc,release",
        xmlName: "IntelliJ IDEA",
        xmlChannels: ["IC-IU-EAP-licensing-RELEASE", "IC-IU-EAP-licensing-EAP", "IC-IU-RELEASE-licensing-RELEASE"],
        useXml: true,
        gradlePropertiesPath: path.resolve(
            pathToProjectRoot,
            "components/ide/jetbrains/backend-plugin/gradle-latest.properties",
        ),
        gradlePropertiesTemplate: `# Code generated by gha-update-image/index-jb-platform-update.ts. DO NOT EDIT.
# See https://plugins.jetbrains.com/docs/intellij/build-number-ranges.html
# for insight into build numbers and IntelliJ Platform versions.
pluginSinceBuild={{pluginSinceBuild}}
pluginUntilBuild={{pluginUntilBuild}}
# Plugin Verifier integration -> https://github.com/JetBrains/gradle-intellij-plugin#plugin-verifier-dsl
# See https://jb.gg/intellij-platform-builds-list for available build versions.
pluginVerifierIdeVersions={{pluginVerifierIdeVersions}}
# Version from "com.jetbrains.intellij.idea" which can be found at https://www.jetbrains.com/intellij-repository/snapshots
platformVersion={{platformVersion}}
`,
    },
    {
        id: 2,
        taskName: "Latest Frontend Plugin",
        productCode: "GW",
        productType: "eap,rc,release",
        xmlName: "Gateway",
        xmlChannels: ["GW-EAP-licensing-EAP", "GW-RELEASE-licensing-RELEASE", "GW-EAP-licensing-RELEASE"],
        useXml: true,
        gradlePropertiesPath: path.resolve(
            pathToProjectRoot,
            "components/ide/jetbrains/gateway-plugin/gradle-latest.properties",
        ),
        gradlePropertiesTemplate: `# Code generated by gha-update-image/index-jb-platform-update.ts. DO NOT EDIT.
# See https://plugins.jetbrains.com/docs/intellij/build-number-ranges.html
# for insight into build numbers and IntelliJ Platform versions.
pluginSinceBuild={{pluginSinceBuild}}
pluginUntilBuild={{pluginUntilBuild}}
# Plugin Verifier integration -> https://github.com/JetBrains/gradle-intellij-plugin#plugin-verifier-dsl
# See https://jb.gg/intellij-platform-builds-list for available build versions.
pluginVerifierIdeVersions={{pluginVerifierIdeVersions}}
# Version from "com.jetbrains.gateway" which can be found at https://www.jetbrains.com/intellij-repository/snapshots
platformVersion={{platformVersion}}
`,
    },
    {
        id: 3,
        taskName: "Stable Backend Plugin",
        productCode: "GW",
        productType: "release",
        xmlName: "Gateway",
        xmlChannels: ["GW-RELEASE-licensing-RELEASE"],
        useXml: false,
        gradlePropertiesPath: path.resolve(
            pathToProjectRoot,
            "components/ide/jetbrains/gateway-plugin/gradle-stable.properties",
        ),
        gradlePropertiesTemplate: `# Code generated by gha-update-image/index-jb-platform-update.ts. DO NOT EDIT.
# See https://plugins.jetbrains.com/docs/intellij/build-number-ranges.html
# for insight into build numbers and IntelliJ Platform versions.
pluginSinceBuild={{pluginSinceBuild}}
pluginUntilBuild={{pluginUntilBuild}}
# Plugin Verifier integration -> https://github.com/JetBrains/gradle-intellij-plugin#plugin-verifier-dsl
# See https://jb.gg/intellij-platform-builds-list for available build versions.
pluginVerifierIdeVersions={{pluginVerifierIdeVersions}}
# Version from "com.jetbrains.gateway" which can be found at https://www.jetbrains.com/updates/updates.xml
platformVersion={{platformVersion}}
`,
    },
];

const { values } = parseArgs({
    args: Bun.argv,
    options: {
        task: {
            type: "string",
        },
    },
    strict: true,
    allowPositionals: true,
});

const taskID = Number.parseInt(values.task ?? "NaN");

const target = targetInfo.find((e) => e.id === taskID);
if (!target) {
    throw new Error(
        `Invalid task id: ${taskID}, update cmd with \`--task="<name>"\`, available tasks: \n\t- ${targetInfo
            .map((e) => `(${e.id}) ${e.taskName}`)
            .join("\n\t- ")}`,
    );
}
console.log("Updating", target.taskName);
jbUpdatePlatformVersion(target);
