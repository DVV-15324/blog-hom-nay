CREATE DATABASE bloghomnay;
GO

USE bloghomnay;
GO

CREATE TABLE users (
    id INT IDENTITY(1,1) PRIMARY KEY,
    email NVARCHAR(255) NOT NULL,
    last_name NVARCHAR(50),
    first_name NVARCHAR(50),
    avatar NVARCHAR(255),
    phone NVARCHAR(20), 
    address NVARCHAR(255),
    deleted_at DATETIME DEFAULT GETDATE(),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE()
);
GO

CREATE TABLE auths (
    id INT IDENTITY(1,1) PRIMARY KEY,
    email NVARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    password NVARCHAR(255) NOT NULL,
    salt NVARCHAR(255) NOT NULL,
	auth_type NVARCHAR(50) NOT NULL DEFAULT 'local',
    deleted_at DATETIME DEFAULT GETDATE(),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
GO

CREATE TABLE categories (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(255) NOT NULL UNIQUE,
<<<<<<< HEAD
    img NVARCHAR(255) NOT NULL UNIQUE,
    description NVARCHAR(MAX),
=======
    description TEXT,
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
    updated_at DATETIME DEFAULT GETDATE(),
    created_at DATETIME DEFAULT GETDATE(),
    deleted_at DATETIME DEFAULT GETDATE()
);
GO

CREATE TABLE posts (
    id INT IDENTITY(1,1) PRIMARY KEY,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    title NVARCHAR(255) NOT NULL,
    content NVARCHAR(MAX),
<<<<<<< HEAD
    description NVARCHAR(MAX),
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
    updated_at DATETIME DEFAULT GETDATE(),
    created_at DATETIME DEFAULT GETDATE(),
    deleted_at DATETIME DEFAULT GETDATE(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
GO

CREATE TABLE tags (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(255) UNIQUE NOT NULL
);
GO

CREATE TABLE post_tags (
    id INT IDENTITY(1,1) PRIMARY KEY,
<<<<<<< HEAD
    post_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    CONSTRAINT UQ_post_tag UNIQUE (post_id, tag_id)
);

=======
    post_id INT,  -- Nếu không ghi gì thì mặc định là có thể NULL
    tag_id INT,   -- Cũng vậy, nếu không ghi gì thì mặc định có thể NULL
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
GO
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8

CREATE TABLE comments (
    id INT IDENTITY(1,1) PRIMARY KEY,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
<<<<<<< HEAD
    content NVARCHAR(MAX) NOT NULL,
=======
    content TEXT NOT NULL,
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
    deleted_at DATETIME DEFAULT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
GO
<<<<<<< HEAD

CREATE TABLE img (
    id INT IDENTITY(1,1) PRIMARY KEY,
    user_id INT NOT NULL,
    img VARCHAR(max),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
GO

CREATE TABLE post_likes (
    id INT IDENTITY(1,1) PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),

    CONSTRAINT UQ_post_likes_user_post UNIQUE (user_id, post_id),

    CONSTRAINT FK_post_likes_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT FK_post_likes_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
