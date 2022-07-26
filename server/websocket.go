/* Copyright 2022 David Bulkow

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"basicblog"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func ws(entries []basicblog.BlogEntry) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("websocket upgrade:", err)
			return
		}
		defer c.Close()

		for _, e := range entries {
			buf, err := json.MarshalIndent(e, "", "    ")
			if err != nil {
				log.Println("json marshal", err)
			}

			err = c.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}

		time.Sleep(time.Minute)
	})
}
