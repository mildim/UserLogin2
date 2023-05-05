# /user/register
## Request:
- First Name (a-zA-Z) max 20 
- Last Name (a-zA-Z) max 30 
- Email* (email standard) max 80 
- Password* (alphanumeric + special characters) min 8 max 20 
- Confirmed Password* (alphanumeric + special characters) min 8 max 20 
- Role (user or admin)
## Response:
- response code (HTTP response status code) 
- status code (0 .. n) 
- message (ok .. n) 

# /user/activate
## Request:
- Email*
- Activation Key*
## Response:
- response code (HTTP response status code) 
- status code (0 .. n) 
- message (ok .. n)

# /user/login
## Request:
- Email*
- Password*
## Response:
- response code (HTTP response status code) 
- status code (0 .. n) 
- message (ok .. n) 
- ID (ObjectID)
- First Name
- Last Name
- Password
- Email
- Token
- Refresh Token
- User ID (string)

# Error codes:
- 0-9 - Database errors
- 10-19 - Controlers errors
- 20-29 - Token(Helper) errors
- 30-39 - Middleware errors
