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

// Full details of a specific event
type GetCharactersCharacterIdCalendarEventIdOk struct {

	// date string
	Date time.Time `json:"date,omitempty"`

	// Length in minutes
	Duration int64 `json:"duration,omitempty"`

	// event_id integer
	EventId int32 `json:"event_id,omitempty"`

	// importance integer
	Importance int32 `json:"importance,omitempty"`

	// owner_id integer
	OwnerId int64 `json:"owner_id,omitempty"`

	// owner_name string
	OwnerName string `json:"owner_name,omitempty"`

	// owner_type string
	OwnerType string `json:"owner_type,omitempty"`

	// response string
	Response string `json:"response,omitempty"`

	// text string
	Text string `json:"text,omitempty"`

	// title string
	Title string `json:"title,omitempty"`
}
