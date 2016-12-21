/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.4
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdMail200Ok struct {

	// From whom the mail was sent
	From int32 `json:"from,omitempty"`

	// is_read boolean
	IsRead bool `json:"is_read,omitempty"`

	// labels array
	Labels []int64 `json:"labels,omitempty"`

	// mail_id integer
	MailId int64 `json:"mail_id,omitempty"`

	// Recipients of the mail
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`

	// Mail subject
	Subject string `json:"subject,omitempty"`

	// When the mail was sent
	Timestamp time.Time `json:"timestamp,omitempty"`
}
