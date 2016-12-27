/*
Copyright 2016 Medcl (m AT medcl.net)

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

package pipe

import (. "github.com/medcl/gopa/core/pipeline"
	"github.com/medcl/gopa/core/tasks"
	"github.com/cihub/seelog"
	"github.com/medcl/gopa/core/types"
	"errors"
)

type Start struct {
	ID *string
	Task *types.Task
}

func (this Start) Name() string {
	return "start"
}

func (this Start) Process(context *Context) (*Context, error) {

	seelog.Trace("start process")

	var task *types.Task
	if(this.Task!=nil){
		task=this.Task
	}else if(this.ID !=nil){
		//init task record
		t,err:=tasks.GetTask(*this.ID)
		if(err!=nil){
			panic(err)
		}
		task=&t
	}else{
		panic(errors.New("task not set"))
	}



	context.Set(CONTEXT_CRAWLER_TASK,task)
	context.Set(CONTEXT_ORIGINAL_URL,task.Seed.Url) //TODO remove
	context.Set(CONTEXT_URL,task.Seed.Url)  //TODO remove
	context.Set(CONTEXT_DEPTH,task.Seed.Depth)  //TODO remove
	context.Set(CONTEXT_REFERENCE_URL,task.Seed.Reference)  //TODO remove

	return context, nil
}

