/*
Copyright 2019 vChain, Inc.

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

package server

import (
	"github.com/codenotary/immudb/pkg/db"
	"github.com/codenotary/immudb/pkg/logger"
)

type ImmuServer struct {
	Topic  *db.Topic
	Logger logger.Logger
}

func (s *ImmuServer) Errorf(f string, v ...interface{}) {
	if s.Logger == nil {
		return
	}
	s.Logger.Errorf(f, v...)
}

func (s *ImmuServer) Warningf(f string, v ...interface{}) {
	s.Logger.Warningf(f, v...)
}

func (s *ImmuServer) Infof(f string, v ...interface{}) {
	s.Logger.Infof(f, v...)
}

func (s *ImmuServer) Debugf(f string, v ...interface{}) {
	s.Logger.Debugf(f, v...)
}
