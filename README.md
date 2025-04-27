HaPTIT
CÁC HÀM TRONG REDIS
1. String (set/get) – Lưu trữ đơn giản
   Sao chép mã
   // Ghi dữ liệu
   redisClient.Set(ctx, "key", "value", time.Hour)

// Đọc dữ liệu
val, err := redisClient.Get(ctx, "key").Result()

// Xóa dữ liệu
redisClient.Del(ctx, "key")
🧮 2. Hash – Lưu trữ kiểu object
go
Sao chép mã
// Lưu thông tin user như 1 map
redisClient.HSet(ctx, "user:1", "username", "sonthuy", "email", "a@gmail.com")

// Lấy toàn bộ
result, _ := redisClient.HGetAll(ctx, "user:1").Result()

// Lấy 1 field
email, _ := redisClient.HGet(ctx, "user:1", "email").Result()
🧺 3. List – Danh sách có thứ tự
go
Sao chép mã
// Thêm vào đầu danh sách
redisClient.LPush(ctx, "queue", "item1", "item2")

// Lấy phần tử
val, _ := redisClient.LPop(ctx, "queue").Result()
📋 4. Set – Danh sách không trùng nhau
go
Sao chép mã
redisClient.SAdd(ctx, "tags", "go", "redis", "cache")

// Kiểm tra phần tử có trong set không
exists, _ := redisClient.SIsMember(ctx, "tags", "go").Result()
📊 5. Sorted Set – Có điểm số, dùng cho ranking
go
Sao chép mã
redisClient.ZAdd(ctx, "leaderboard", &redis.Z{Score: 100, Member: "user1"})

// Lấy top
top, _ := redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, 9).Result()
🔄 6. TTL và thời gian sống
go
Sao chép mã
// Xem thời gian sống còn lại
ttl, _ := redisClient.TTL(ctx, "key").Result()

// Cập nhật TTL
redisClient.Expire(ctx, "key", time.Minute*10)
🔍 7. Keys – Dò key (nên tránh dùng nhiều)
go
Sao chép mã
keys, _ := redisClient.Keys(ctx, "user:*").Result() // Dò key (chậm nếu key nhiều)
✨ Một số lệnh nâng cao:
Hàm	Mục đích
Exists	Kiểm tra key có tồn tại
FlushDB	Xóa toàn bộ key trong DB hiện tại
Incr, Decr	Tăng/giảm giá trị số
SetNX	Set key nếu chưa tồn tại

Trong nghiệp vụ logic Thằng A xóa dữ liệu trong database 

Luồng thực tế:
✅ Người dùng đăng nhập → token + quyền được lưu vào Redis

✅ Gọi API → kiểm tra token + quyền trong Redis

❌ Token không có / quyền không đủ → từ chối

🔁 Đổi quyền → xóa cache → lần sau sẽ lấy từ DB & cache lại

