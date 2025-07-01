const express = require("express");
const multer = require("multer");
const fs = require("fs");
const path = require("path");
const cors = require("cors");

const app = express();

const corsOptions = {
  origin: "*",
  methods: ["GET", "POST"],
  allowedHeaders: ["Content-Type"],
};
app.use(cors(corsOptions));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Thư mục lưu ảnh
const UPLOAD_DIR = path.join(__dirname, "uploads");
if (!fs.existsSync(UPLOAD_DIR)) {
  fs.mkdirSync(UPLOAD_DIR);
}

// Mapping uploads/ vào root /
app.use(express.static(UPLOAD_DIR));

app.use((req, res, next) => {
  // Nếu URL có đuôi là .jpg/ hoặc .png/ hoặc .jpeg/... thì redirect về không có dấu /
  if (/\.(jpg|jpeg|png|gif|webp|svg)\/$/.test(req.url)) {
    const newUrl = req.url.replace(/\/+$/, "");
    return res.redirect(301, newUrl);
  }
  next();
});

// Cấu hình multer
const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    cb(null, UPLOAD_DIR);
  },
  filename: function (req, file, cb) {
    const uniqueName = Date.now() + path.extname(file.originalname);
    cb(null, uniqueName);
  },
});
const upload = multer({ storage });

// Upload API
// Upload API
app.post("/upload-image", upload.single("image"), (req, res) => {
  if (!req.file) {
    return res.status(400).json({ error: "No image uploaded" });
  }
  const filename = req.file.filename;
  const fileUrl = `http://image.bloghomnay.com/${filename}`;

  res.json({ img: fileUrl });
});



// Run server
const PORT = 80;
app.listen(PORT, () => {
  console.log(`✅ Server running at http://localhost`);
});
