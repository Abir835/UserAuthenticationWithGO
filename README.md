Features
User Management

    User Registration: Users can register using their email and password.
    User Login: Users can log in using their credentials.
    OTP Verification: OTP is implemented for email verification during registration and for additional security during login.

Authentication and Authorization

    JWT Authentication: JSON Web Tokens are used for user authentication.
    Authorization: Endpoints are secured using JWT tokens. Admin and normal users have different levels of access.

Email Integration

    Email Notifications: Users receive a welcome email upon successful registration. OTP is sent via email for verification during registration and login.

User Roles

    Admin: Can manage (add, update, delete) books.
    Normal User: Can view and purchase books.

Books Management (Admin)

    Create Book: Admin users can add new books with details like title, author, price, and description.
    Update Book: Admin users can update existing book details.
    Delete Book: Admin users can delete books from the store inventory.

Books Viewing and Purchasing (Normal User)

    View Books: Normal users can view the list of available books.
    View Book Details: Normal users can see detailed information about a specific book.
    Purchase Books: Implement functionality for normal users to purchase books.
    
Books Purchasing (Normal User)

    Purchase Books: Now i implement purchase book.
    Price and Pcs and Toal Price we can validate.


Database

    We are Using Mysql Data
    1. Create Database Instance like "bookstore"
    2. Then how to configure database
    Then how to configure database 
    dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbName)

	db, err := gorm.Open("mysql", dsn)

Dependencies

    github.com/dgrijalva/jwt-go v3.2.0
    github.com/gorilla/mux v1.8.1
    github.com/jinzhu/gorm v1.9.16
    github.com/joho/godotenv v1.5.1
    golang.org/x/crypto v0.25.0
    gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df

Setup Instructions

    Clone the Repository: git clone <repository-url>
    Install Dependencies: go mod tidy
    Set Environment Variables:
        Create a .env file and set variables for database connection, SMTP configuration, etc.
    Run the Application: go run cmd/main.go
    Access the Application: Open your browser and go to http://localhost:8000

API Endpoints
Authentication

    POST /register: Register a new user.
    POST /login: Authenticate user and generate JWT token.
    POST /verify-otp: Verify OTP during registration and login.
    POST /logout: Invalidate user session and logout.

Admin Operations

    GET /admin/dashboard: View admin dashboard.
    POST /admin/books: Add a new book.
    PUT /admin/books/{id}: Update book details.
    DELETE /admin/books/{id}: Delete a book.
    GET /admin/books: Get list of all books.
    GET /admin/books/{id}: Get details of a specific book.

User Operations

    GET /user/profile: View user profile.
    GET /user/books: Get list of available books for normal users.
    GET /user/books/{id}: Get details of a specific book for normal users
    POST /user/purchase: Book Purchase