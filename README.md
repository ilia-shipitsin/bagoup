bagoup
======
*(pronounced BAAGoop)*

A backup utility for Mac OS Messages, implemented in Go, inspired by
[Baskup](http://peterkaminski09.github.io/baskup/).

# Usage

1. Clone this repo into your GOPATH `git clone git@github.com:tagatac/bagoup.git`.
1. Open Finder.
1. Navigate to **~/Library/Messages**.
1. Right-click on **chat.db**, and click **Copy "chat.db"** in the context menu.
1. Navigate to your clone of this repo.
1. Right-click in the `bagoup` directory, and click **Paste Item** in the
context menu.
1. `go run bagoup.go`

# License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

# Author

Copyright (C) 2020 [David Tagatac](david@tagatac.net)
