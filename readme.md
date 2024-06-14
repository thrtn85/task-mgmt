# Title: Task Management Web App

### Brief:
Create a web application for managing tasks. The app should allow users to register, log in, and manage their tasks. Each task should have a title, description, due date, and status (e.g., pending, completed). 

### Key Features:
1. User Authentication: Users should be able to register for an account and log in securely.
2. Task Management: Once logged in, users should be able to create, view, edit, and delete tasks.
3. Task Sorting and Filtering: Users should have options to sort tasks by due date, status, or title. Additionally, they should be able to filter tasks based on their status (e.g., view only pending tasks).
4. User Profile: Users should have a profile page where they can update their information, such as username and password.
5. Responsive Design: Ensure the web app is responsive and works well on both desktop and mobile devices.

### Optional Features:
1. Task Categories: Allow users to categorize tasks into different groups (e.g., work, personal).
2. Task Reminders: Implement a feature that sends email reminders to users for upcoming tasks.
3. Collaboration: Enable users to share tasks or task lists with other registered users.
4. Task Comments: Allow users to add comments or notes to tasks.
5. Task Attachments: Enable users to upload and attach files or documents to tasks.

### Tech Stack:
- Backend: Go (Golang) with a web framework like Gin or Echo.
- Frontend: HTML, CSS, JavaScript (you can use a frontend framework like React or Vue.js if you're comfortable).
- Database: Use PostgreSQL or MySQL for storing user accounts and tasks. Sqlite for dev.
- Authentication: Implement JWT (JSON Web Tokens) for user authentication.
- Deployment: Deploy the web app to a cloud platform like Heroku or AWS.