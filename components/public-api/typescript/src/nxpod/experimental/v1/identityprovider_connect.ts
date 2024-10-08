/**
 * Copyright (c) 2024 Nxpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file nxpod/experimental/v1/identityprovider.proto (package nxpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetIDTokenRequest, GetIDTokenResponse } from "./identityprovider_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service nxpod.experimental.v1.IdentityProviderService
 */
export const IdentityProviderService = {
  typeName: "nxpod.experimental.v1.IdentityProviderService",
  methods: {
    /**
     * GetIDToken produces a new OIDC ID token (https://openid.net/specs/openid-connect-core-1_0.html#ImplicitIDToken)
     *
     * @generated from rpc nxpod.experimental.v1.IdentityProviderService.GetIDToken
     */
    getIDToken: {
      name: "GetIDToken",
      I: GetIDTokenRequest,
      O: GetIDTokenResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
