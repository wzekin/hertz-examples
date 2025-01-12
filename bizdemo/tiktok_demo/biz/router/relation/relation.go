/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by hertz generator. DO NOT EDIT.

package relation

import (
	relation "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/handler/relation"
	"github.com/cloudwego/hertz/pkg/app/server"
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
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action := _relation.Group("/action", _actionMw()...)
				_action.POST("/", append(_relation_ctionMw(), relation.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list := _follow.Group("/list", _listMw()...)
					_list.GET("/", append(_relationfollowlistMw(), relation.RelationFollowList)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list0 := _follower.Group("/list", _list0Mw()...)
					_list0.GET("/", append(_relationfollowerlistMw(), relation.RelationFollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list1 := _friend.Group("/list", _list1Mw()...)
					_list1.GET("/", append(_relationfriendlistMw(), relation.RelationFriendList)...)
				}
			}
		}
	}
}
