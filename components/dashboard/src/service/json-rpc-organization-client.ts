/**
 * Copyright (c) 2023 Nxpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { PartialMessage } from "@bufbuild/protobuf";
import { CallOptions, PromiseClient } from "@connectrpc/connect";
import { OrganizationService } from "@nxpod/public-api/lib/nxpod/v1/organization_connect";
import {
    CreateOrganizationRequest,
    CreateOrganizationResponse,
    DeleteOrganizationMemberRequest,
    DeleteOrganizationMemberResponse,
    DeleteOrganizationRequest,
    DeleteOrganizationResponse,
    GetOrganizationInvitationRequest,
    GetOrganizationInvitationResponse,
    GetOrganizationRequest,
    GetOrganizationResponse,
    GetOrganizationSettingsRequest,
    GetOrganizationSettingsResponse,
    JoinOrganizationRequest,
    JoinOrganizationResponse,
    ListOrganizationMembersRequest,
    ListOrganizationMembersResponse,
    ListOrganizationWorkspaceClassesRequest,
    ListOrganizationWorkspaceClassesResponse,
    ListOrganizationsRequest,
    ListOrganizationsResponse,
    OrganizationSettings,
    ResetOrganizationInvitationRequest,
    ResetOrganizationInvitationResponse,
    UpdateOrganizationMemberRequest,
    UpdateOrganizationMemberResponse,
    UpdateOrganizationRequest,
    UpdateOrganizationResponse,
    UpdateOrganizationSettingsRequest,
    UpdateOrganizationSettingsResponse,
} from "@nxpod/public-api/lib/nxpod/v1/organization_pb";
import { getNxpodService } from "./service";
import { converter } from "./public-api";
import { ApplicationError, ErrorCodes } from "@nxpod/nxpod-protocol/lib/messaging/error";
import { OrgMemberRole } from "@nxpod/nxpod-protocol";

export class JsonRpcOrganizationClient implements PromiseClient<typeof OrganizationService> {
    async createOrganization(
        request: PartialMessage<CreateOrganizationRequest>,
        options?: CallOptions | undefined,
    ): Promise<CreateOrganizationResponse> {
        if (!request.name) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "name is required");
        }
        const result = await getNxpodService().server.createTeam(request.name);
        return new CreateOrganizationResponse({
            organization: converter.toOrganization(result),
        });
    }

    async listOrganizationWorkspaceClasses(
        request: PartialMessage<ListOrganizationWorkspaceClassesRequest>,
        options?: CallOptions | undefined,
    ): Promise<ListOrganizationWorkspaceClassesResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const list = await getNxpodService().server.getOrgWorkspaceClasses(request.organizationId);
        return new ListOrganizationWorkspaceClassesResponse({
            workspaceClasses: list.map((e) => converter.toWorkspaceClass(e)),
        });
    }

    async getOrganization(
        request: PartialMessage<GetOrganizationRequest>,
        options?: CallOptions | undefined,
    ): Promise<GetOrganizationResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "id is required");
        }
        const result = await getNxpodService().server.getTeam(request.organizationId);

        return new GetOrganizationResponse({
            organization: converter.toOrganization(result),
        });
    }

    async updateOrganization(
        request: PartialMessage<UpdateOrganizationRequest>,
        options?: CallOptions | undefined,
    ): Promise<UpdateOrganizationResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        if (!request.name) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "name is required");
        }
        await getNxpodService().server.updateTeam(request.organizationId, {
            name: request.name,
        });
        return new UpdateOrganizationResponse();
    }

    async listOrganizations(
        request: PartialMessage<ListOrganizationsRequest>,
        options?: CallOptions | undefined,
    ): Promise<ListOrganizationsResponse> {
        const result = await getNxpodService().server.getTeams();
        return new ListOrganizationsResponse({
            organizations: result.map((team) => converter.toOrganization(team)),
        });
    }

    async deleteOrganization(
        request: PartialMessage<DeleteOrganizationRequest>,
        options?: CallOptions | undefined,
    ): Promise<DeleteOrganizationResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        await getNxpodService().server.deleteTeam(request.organizationId);
        return new DeleteOrganizationResponse();
    }

    async getOrganizationInvitation(
        request: PartialMessage<GetOrganizationInvitationRequest>,
        options?: CallOptions | undefined,
    ): Promise<GetOrganizationInvitationResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const result = await getNxpodService().server.getGenericInvite(request.organizationId);
        return new GetOrganizationInvitationResponse({
            invitationId: result.id,
        });
    }

    async joinOrganization(
        request: PartialMessage<JoinOrganizationRequest>,
        options?: CallOptions | undefined,
    ): Promise<JoinOrganizationResponse> {
        if (!request.invitationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "invitationId is required");
        }
        const result = await getNxpodService().server.joinTeam(request.invitationId);
        return new JoinOrganizationResponse({
            organizationId: result.id,
        });
    }

    async resetOrganizationInvitation(
        request: PartialMessage<ResetOrganizationInvitationRequest>,
        options?: CallOptions | undefined,
    ): Promise<ResetOrganizationInvitationResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const newInvite = await getNxpodService().server.resetGenericInvite(request.organizationId);
        return new ResetOrganizationInvitationResponse({
            invitationId: newInvite.id,
        });
    }

    async listOrganizationMembers(
        request: PartialMessage<ListOrganizationMembersRequest>,
        options?: CallOptions | undefined,
    ): Promise<ListOrganizationMembersResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const result = await getNxpodService().server.getTeamMembers(request.organizationId);
        return new ListOrganizationMembersResponse({
            members: result.map((member) => converter.toOrganizationMember(member)),
        });
    }

    async updateOrganizationMember(
        request: PartialMessage<UpdateOrganizationMemberRequest>,
        options?: CallOptions | undefined,
    ): Promise<UpdateOrganizationMemberResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        if (!request.userId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "userId is required");
        }
        if (!request.role) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "role is required");
        }
        await getNxpodService().server.setTeamMemberRole(
            request.organizationId,
            request.userId,
            converter.fromOrgMemberRole(request.role),
        );
        return new UpdateOrganizationMemberResponse();
    }

    async deleteOrganizationMember(
        request: PartialMessage<DeleteOrganizationMemberRequest>,
        options?: CallOptions | undefined,
    ): Promise<DeleteOrganizationMemberResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        if (!request.userId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "userId is required");
        }
        await getNxpodService().server.removeTeamMember(request.organizationId, request.userId);
        return new DeleteOrganizationMemberResponse();
    }

    async getOrganizationSettings(
        request: PartialMessage<GetOrganizationSettingsRequest>,
        options?: CallOptions | undefined,
    ): Promise<GetOrganizationSettingsResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const result = await getNxpodService().server.getOrgSettings(request.organizationId);
        return new GetOrganizationSettingsResponse({
            settings: converter.toOrganizationSettings(result),
        });
    }

    async updateOrganizationSettings(
        request: PartialMessage<UpdateOrganizationSettingsRequest>,
        options?: CallOptions | undefined,
    ): Promise<UpdateOrganizationSettingsResponse> {
        if (!request.organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const update: Partial<OrganizationSettings> = {
            workspaceSharingDisabled: request?.workspaceSharingDisabled,
            defaultWorkspaceImage: request?.defaultWorkspaceImage,
            allowedWorkspaceClasses: request?.allowedWorkspaceClasses,
            restrictedEditorNames: request?.restrictedEditorNames,
            defaultRole: request?.defaultRole,
        };
        if (request.updatePinnedEditorVersions) {
            update.pinnedEditorVersions = request.pinnedEditorVersions;
        } else if (request.pinnedEditorVersions && Object.keys(request.pinnedEditorVersions).length > 0) {
            throw new ApplicationError(
                ErrorCodes.BAD_REQUEST,
                "updatePinnedEditorVersions is required to be true to update pinnedEditorVersions",
            );
        }
        if (request.updateRestrictedEditorNames) {
            update.restrictedEditorNames = request.restrictedEditorNames;
        } else if (request.restrictedEditorNames && request.restrictedEditorNames.length > 0) {
            throw new ApplicationError(
                ErrorCodes.BAD_REQUEST,
                "updateRestrictedEditorNames is required to be true to update restrictedEditorNames",
            );
        }
        await getNxpodService().server.updateOrgSettings(request.organizationId, {
            ...update,
            defaultRole: request.defaultRole as OrgMemberRole,
            timeoutSettings: {
                inactivity: converter.toDurationString(request.timeoutSettings?.inactivity),
                denyUserTimeouts: request.timeoutSettings?.denyUserTimeouts,
            },
        });
        return new UpdateOrganizationSettingsResponse();
    }
}
