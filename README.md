# chirp-core-service

![Go](https://img.shields.io/badge/Go-blue?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-lightgreen?style=for-the-badge&logo=gin&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/Postgres-yellow?style=for-the-badge&logo=postgresql&logoColor=white)


Core functionality - posts, comments and follow logic for chirp

Built with Go, Gin and PostgreSQL 


## API Endpoints

BASE_URL: `localhost:8002` (if running locally; unless different port used)

### Root

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/`

Authentication: Public

Returns `"Hello from  chirp-core-service"`

Response 
```
{
    "message": "Hello from chirp-core-service"
}
```

### Fetch All Posts

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/posts/`

Authentication: Public

Retrieves all posts; latest first

Response 
```
{
    "posts": [
        {
            "id": 7,
            "user_id": 17,
            "content": "alpha",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-27T21:27:06.018375Z"
        },
        {
            "id": 6,
            "user_id": 17,
            "content": "beta",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-27T21:26:12.814793Z"
        },
        {
            "id": 5,
            "user_id": 17,
            "content": "beta",
            "likes_count": 1,
            "comments_count": 4,
            "created_at": "2025-10-21T15:29:52.901076Z"
        },
        {
            "id": 3,
            "user_id": 17,
            "content": "post by user 2",
            "likes_count": 0,
            "comments_count": 1,
            "created_at": "2025-10-21T15:29:37.276353Z"
        },
        {
            "id": 2,
            "user_id": 16,
            "content": "post 2\n 2nd post",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-21T14:49:27.282497Z"
        },
        {
            "id": 1,
            "user_id": 16,
            "content": "post1",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-21T14:48:26.413072Z"
        }
    ]
}
```

###  Fetch Posts by User

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/posts/user/:id`

Authentication: Public

Retrieves all posts; latest first; for user id provided in params

Response 
```
{
    "posts": [
        {
            "id": 7,
            "user_id": 17,
            "content": "alpha",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-27T21:27:06.018375Z"
        },
        {
            "id": 6,
            "user_id": 17,
            "content": "beta",
            "likes_count": 0,
            "comments_count": 0,
            "created_at": "2025-10-27T21:26:12.814793Z"
        },
        {
            "id": 5,
            "user_id": 17,
            "content": "beta",
            "likes_count": 1,
            "comments_count": 4,
            "created_at": "2025-10-21T15:29:52.901076Z"
        },
        {
            "id": 3,
            "user_id": 17,
            "content": "post by user 2",
            "likes_count": 0,
            "comments_count": 1,
            "created_at": "2025-10-21T15:29:37.276353Z"
        }
    ]
}
```

### Create Post

![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `/posts`

Authentication: JWT to be provided in `x-jwt-token` header

Creates post for user associated with provided JWT

Request 
```
{
    "content": "sample post"
}
```

Response 
```
{
    "content": "sample post",
    "created_at": "2025-10-28T20:14:50.959313573+05:30",
    "message": "Created post",
    "post_id": 8
}
```

### Get Post by Id 

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/posts/:id`

Authentication: Public

Retrieves post associated with provided id

Response 
```
{
    "post": {
        "id": 5,
        "user_id": 17,
        "content": "beta",
        "likes_count": 1,
        "comments_count": 4,
        "created_at": "2025-10-21T15:29:52.901076Z"
    }
}
```

### Delete Own Post

![DELETE](https://img.shields.io/badge/DELETE-%23E74C3C?style=for-the-badge&logo=postman&logoColor=white)  `/posts/:id`

Authentication: JWT to be provided under `x-jwt-token` header

Deletes post with given id if associated with provided user's JWT

Response
```
{
    "message": "Post deleted successfully"
}
```

###  Like Post

![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `/posts/:id/like`

Authentication: JWT to be provided in `x-jwt-token` header

Likes post with provided id

Response 
```
{
    "message": "Post liked"
}
```

###  Unlike Post

![DELETE](https://img.shields.io/badge/DELETE-%23E74C3C?style=for-the-badge&logo=postman&logoColor=white) `/posts/:id/unlike`

Authentication: JWT to be provided in `x-jwt-token` header

Unlikes post with provided id

Response 
```
{
    "message": "Post unliked"
}
```

### Fetch Comments by post

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/comments/post/:post_id`

Authentication: Public

Retrieves comments associated with provided post_id

Response
```
{
    "comments": [
        {
            "id": 8,
            "post_id": 5,
            "user_id": 17,
            "content": "random comment",
            "likes_count": 1,
            "created_at": "2025-10-22T14:20:32.938303Z"
        },
        {
            "id": 3,
            "post_id": 5,
            "user_id": 17,
            "content": "neutral comment",
            "likes_count": 0,
            "created_at": "2025-10-22T14:18:42.245114Z"
        },
        {
            "id": 2,
            "post_id": 5,
            "user_id": 17,
            "content": "good comment",
            "likes_count": 1,
            "created_at": "2025-10-22T14:18:36.882457Z"
        },
        {
            "id": 1,
            "post_id": 5,
            "user_id": 17,
            "content": "bad comment",
            "likes_count": 0,
            "created_at": "2025-10-22T14:18:31.206377Z"
        }
    ]
}
```

### Create Comment

![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `/comments/post/:post_id`

Authentication: JWT to be provided under `x-jwt-token` header

Post comment under post associated with provided post_id

Request
```
{
    "content": "sample comment"
}
```

Response
```
{
    "comment_id": 9,
    "content": "sample comment",
    "created_at": "2025-10-28T20:33:47.454861921+05:30",
    "message": "Comment created"
}
```

### Delete Own Comment

![DELETE](https://img.shields.io/badge/DELETE-%23E74C3C?style=for-the-badge&logo=postman&logoColor=white)  `/comments/:id`

Authentication: JWT to be provided under `x-jwt-token` header

Deletes comment with given id if associated with provided user's JWT

Response
```
{
    "message": "Comment deleted successfully"
}
```

###  Like Comment

![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `/comments/:id/like`

Authentication: JWT to be provided in `x-jwt-token` header

Likes comment with provided id

Response 
```
{
    "message": "Comment liked"
}
```

###  Unlike Comment

![DELETE](https://img.shields.io/badge/DELETE-%23E74C3C?style=for-the-badge&logo=postman&logoColor=white) `/comments/:id/unlike`

Authentication: JWT to be provided in `x-jwt-token` header

Unlikes comment with provided id

Response 
```
{
    "message": "Comment unliked"
}
```

### Follow User

![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `/follow/:id`

Authentication: JWT to be provided in `x-jwt-token` header

User associated with provided JWT token follows user with provided id

Response
```
{
    "message": "Followed user"
}
```

### Unfollow User

![DELETE](https://img.shields.io/badge/DELETE-%23E74C3C?style=for-the-badge&logo=postman&logoColor=white) `/follow/:id`

Authentication: JWT to be provided in `x-jwt-token` header

User associated with provided JWT token unfollows user with provided id

Response
```
{
    "message": "Unfollowed user"
}
```

### Get Followers

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/follow/followers/:id`

Authentication: Public

Get followers for user associated with provided id

Response
```
{
    "followers": [
        {
            "id": 17,
            "username": "test2",
            "email": "test2@gmail.com",
            "bio": "",
            "likes_count": 0,
            "followers_count": 0,
            "following_count": 1,
            "tweets_count": 0,
            "created_at": "2025-10-21T15:28:52.345535Z"
        },
        {
            "id": 18,
            "username": "test3",
            "email": "test3@gmail.com",
            "bio": "",
            "likes_count": 0,
            "followers_count": 0,
            "following_count": 2,
            "tweets_count": 0,
            "created_at": "2025-10-23T20:03:12.704324Z"
        }
    ]
}
```

### Get Following

![GET](https://img.shields.io/badge/GET-%2334D058?style=for-the-badge&logo=postman&logoColor=white) `/follow/following/:id`

Authentication: Public

Get users followed by user associated with provided id

Response
```
{
    "following": [
        {
            "id": 16,
            "username": "test",
            "email": "test@gmail.com",
            "bio": "",
            "likes_count": 0,
            "followers_count": 1,
            "following_count": 0,
            "tweets_count": 0,
            "created_at": "2025-10-21T14:45:58.756466Z"
        },
        {
            "id": 14,
            "username": "root",
            "email": "root@gmail.com",
            "bio": "Alison im lost",
            "likes_count": 0,
            "followers_count": 2,
            "following_count": 0,
            "tweets_count": 0,
            "created_at": "2025-10-18T15:29:04.796502Z"
        }
    ]
}
```

## Project Structure

```
.
├── cmd/auth/           # main function
├── internal/
│   ├── controllers/    # request handlers
│   ├── db/             # database connection and table creation
│   ├── middleware/     # JWT authentication middleware
│   ├── models/         # database model structs 
│   └── utils/          # JWT creation and validation
├── .env         
├── go.mod              # go module declaration and dependencies
├── go.sum              
├── Makefile            # build and run commands
└── README.md
```