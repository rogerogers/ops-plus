// Code generated by hertz generator. DO NOT EDIT.

package Repository

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	repository "github.com/rogerogers/ops-plus/biz/handler/repository"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_repo := _api.Group("/repo", _repoMw()...)
			_repo.GET("/list", append(_listMw(), repository.List)...)
		}
	}
}
