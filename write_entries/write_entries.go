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
	"os"
	"time"

	"basicblog"
)

func main() {
	entries := []basicblog.BlogEntry{
		{
			Created: time.Now(),
			Expires: time.Now().Add(time.Hour),
			Author:  "Some Person",
			Content: "this is a blog entry",
		},
		{
			Created: time.Now(),
			Expires: time.Now().Add(2 * time.Hour),
			Author:  "Some Other Person",
			Content: "this is also a blog entry",
		},
	}

	buf, err := json.MarshalIndent(&entries, "", "    ")
	if err != nil {
		log.Fatal("json marshal", err)
	}

	err = os.WriteFile("blogentries.json", buf, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
