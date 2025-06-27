USE bloghomnay;
GO

INSERT INTO tags (name) VALUES
('GoLang'),
('Docker'),
('Kubernetes'),
('DevOps'),
('Java'),
('Linux'),
('Ubuntu'),
('AWS');
GO

USE bloghomnay;


INSERT INTO categories (name, description, updated_at, created_at, deleted_at, img, tag_id) VALUES
('GoLang', N'Go là một ngôn ngữ lập trình mới do Google thiết kế và phát triển. Nó được kỳ vọng sẽ giúp ngành công nghiệp phần mềm khai thác tối đa nền tảng đa lõi của bộ vi xử lý và hoạt động đa nhiệm tốt hơn.', '2025-06-09 00:05:56.907', '2025-06-09 00:05:56.907', '2025-06-09 00:05:56.907', 'http://image.bloghomnay.com/img_golang.jpg', 1),
('Docker', N'Docker là một dự án mã nguồn mở giúp tự động triển khai các ứng dụng Linux và Windows vào trong các container ảo hóa.', '2025-06-09 00:06:47.590', '2025-06-09 00:06:47.590', '2025-06-09 00:06:47.590', 'http://image.bloghomnay.com/img_docker.png', 2),
('Kubernetes', N'Kubernetes là một nền tảng nguồn mở, khả chuyển, có thể mở rộng để quản lý các ứng dụng được đóng gói và các service, giúp thuận lợi trong việc cấu hình và tự động hoá việc triển khai ứng dụng. Kubernetes là một hệ sinh thái lớn và phát triển nhanh chóng.', '2025-06-09 00:13:35.940', '2025-06-09 00:13:35.940', '2025-06-09 00:13:35.940', 'http://image.bloghomnay.com/img_k8s.png', 3),
('DevOps', N'DevOps (kết hợp của cụm từ tiếng Anh "software DEVelopment" và "information technology OPerationS") là một thuật ngữ để chỉ một tập hợp các hành động trong đó nhấn mạnh sự hợp tác và trao đổi thông tin của các lập trình viên và chuyên viên tin học', '2025-06-09 00:14:26.960', '2025-06-09 00:14:26.960', '2025-06-09 00:14:26.960', 'http://image.bloghomnay.com/img_devops.png', 4),
('Java', N'Java (phiên âm Tiếng Việt: "Gia-va") là một ngôn ngữ lập trình hướng đối tượng, dựa trên lớp được thiết kế để có càng ít phụ thuộc thực thi càng tốt', '2025-06-09 01:15:02.863', '2025-06-09 01:15:02.863', '2025-06-09 01:15:02.863', 'http://image.bloghomnay.com/img_java.png', 5),
('Linux', N'Linux (/ˈlɪnʊks/)[15] là một họ các hệ điều hành tự do nguồn mở tương tự Unix dựa trên nhân Linux,[16] một hạt nhân hệ điều hành được phát hành lần đầu tiên vào ngày 17 tháng 9 năm 1991 bởi Linus Torvalds.', '2025-06-09 01:16:24.617', '2025-06-09 01:16:24.617', '2025-06-09 01:16:24.617', 'http://image.bloghomnay.com/img_linux.jpg', 6),
('Ubuntu', N'Ubuntu (/ʊˈbʊntuː/) là một hệ điều hành máy tính dựa trên Debian GNU/Linux, một bản phân phối Linux thông dụng. Tên của nó bắt nguồn từ "ubuntu" trong tiếng Zulu, có nghĩa là "tình người", mô tả triết lý ubuntu: "Tôi được là chính mình nhờ có những người xung quanh," một khía cạnh tích cực của cộng đồng.', '2025-06-09 01:17:35.867', '2025-06-09 01:17:35.867', '2025-06-09 01:17:35.867', 'http://image.bloghomnay.com/img_ubuntu.png', 7),
('AWS', N'Amazon Web Services (AWS) là một công ty con của Amazon cung cấp các nền tảng điện toán đám mây theo yêu cầu cho các cá nhân, công ty và chính phủ, trên cơ sở trả tiền theo nhu cầu sử dụng (pay-as-you-go).', '2025-06-09 01:18:26.303', '2025-06-09 01:18:26.303', '2025-06-09 01:18:26.303', 'http://image.bloghomnay.com/img_aws.png', 8);

