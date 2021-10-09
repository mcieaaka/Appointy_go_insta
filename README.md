# Go API for Insta Clone

>Clone Repo

>go build

>./Appointy_go_insta

## Instructions

>Open Postman at localhost:8000

>/users POST
- Raw Json Body for users
    <pre>
    {
        "name": "abc",
        "email": "abc@example.com",
        "password": "afaafafa"
    }
    </pre>

>/users/{id} GET

>/posts POST
- Raw JSON Body for posts
    <pre>
    {
        "Caption":"abc",
        "ImageUrl":"%URL%",
        "UserID" :"any user id"
    }
    </pre>

>/posts/{id} GET

>/posts/users/{id} GET