# Snippetbox - A Golang Web Application  

Welcome to **Snippetbox**, a web application for storing and managing snippets of text. This project is built as part of my journey to learn web development in Go (Golang), based on the book *[Let's Go!](https://lets-go.alexedwards.net)* by Alex Edwards.  

## 🚀 Project Overview  

Snippetbox is a simple yet powerful web application that allows users to:  
- Create snippets of text.  
- View snippets by their unique ID.  
- Browse a list of the latest snippets.  

This project is a learning exercise designed to implement the best practices in Go web development, including working with templates, routing, forms, authentication, and more.  

## 🌟 Features  

- **Fast and Efficient**: Built using Go, a statically-typed, compiled language optimized for performance.  
- **Clean Architecture**: Implements the layered design principles from *Let's Go!*  
- **HTML Templates**: Dynamically render web pages using Go's `html/template` package.  
- **Secure**: Includes best practices for input validation and protection against common vulnerabilities like XSS.  
- **PostgreSQL Integration**: Uses PostgreSQL as the database to store snippet information.  

## 📚 What I Learned  

While building this project, I focused on:  
1. Structuring a Go web application for maintainability and scalability.  
2. Using Go's built-in libraries for routing and templating.  
3. Handling forms and managing user input securely.  
4. Working with a database (PostgreSQL) in Go.  
5. Adding session-based authentication.  

## 🔧 Getting Started  

### Prerequisites  
- [Go](https://golang.org/dl/) (1.20 or higher recommended)  
- [PostgreSQL](https://www.postgresql.org/download/)  

### Installation  
1. Clone the repository:  
   ```bash  
   git clone https://github.com/your-username/snippetbox.git  
   cd snippetbox  
   ```  

2. Install dependencies:  
   ```bash  
   go mod tidy  
   ```  

3. Set up the database:  
   - Create a PostgreSQL database.  
   - Apply the SQL schema in `sql/schema.sql`.  
     ```bash  
     psql -U your-username -d snippetbox -f sql/schema.sql  
     ```  

4. Create a `.env` file in the root directory with the following:  
   ```env  
   DB_DSN="user=your-username password=your-password dbname=snippetbox sslmode=disable"  
   ```  

5. Run the application:  
   ```bash  
   go run ./cmd/web  
   ```  

6. Open your browser and visit: [http://localhost:4000](http://localhost:4000)  

## 📂 Project Structure  

```  
snippetbox/  
├── cmd/  
│   └── web/         # Entry point for the application  
├── internal/  
│   ├── models/      # Database models and queries  
│   ├── handlers/    # HTTP handlers for routing  
│   ├── templates/   # HTML templates  
│   ├── middleware/  # Custom middleware functions  
│   └── …  
├── sql/             # Database schema and migrations  
├── go.mod           # Dependencies file  
└── README.md        # Project documentation  
```  

## 🤝 Contributing  

Contributions are welcome! Feel free to submit issues or pull requests to improve this project.  

## 📝 License  

This project is for educational purposes and is not intended for production use.  

## 🙌 Acknowledgments  

- Special thanks to Alex Edwards for writing *Let's Go!*, which served as the primary resource for this project.  

---

Let me know if you'd like to add specific details, like badges, screenshots, or deployment instructions!