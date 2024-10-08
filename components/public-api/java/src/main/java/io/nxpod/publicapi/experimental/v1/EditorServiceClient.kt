// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: nxpod/experimental/v1/editor_service.proto
//
package io.nxpod.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class EditorServiceClient(
  private val client: ProtocolClientInterface,
) : EditorServiceClientInterface {
  override suspend fun listEditorOptions(request: EditorServiceOuterClass.ListEditorOptionsRequest,
      headers: Headers): ResponseMessage<EditorServiceOuterClass.ListEditorOptionsResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "nxpod.experimental.v1.EditorService/ListEditorOptions",
      io.nxpod.publicapi.experimental.v1.EditorServiceOuterClass.ListEditorOptionsRequest::class,
      io.nxpod.publicapi.experimental.v1.EditorServiceOuterClass.ListEditorOptionsResponse::class,
      StreamType.UNARY,
    ),
  )

}
