/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
)

func getPage(host string, port int) (string, error) {
	url := fmt.Sprintf("%s:%d", host, port)
	client, err := redis.Dial("tcp", url)
	if err != nil {
		log.Error("Unable to connect to Redis: %v", err)
		return "", err
	}
	defer client.Close()

	resp, err := client.Cmd("INFO").Str()
	if err != nil {
		log.Error("Unable to retrieve INFO: %v", err)
		return "", err
	}

	return resp, nil
}
