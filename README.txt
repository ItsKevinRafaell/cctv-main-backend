# CCTV AI Pendeteksi Anomali

Proyek ini merupakan sistem **CCTV berbasis AI** yang mampu mendeteksi aktivitas atau perilaku anomali secara otomatis menggunakan *machine learning* dan *computer vision*. Sistem ini dirancang untuk membantu pemantauan keamanan secara real-time dengan tingkat akurasi yang tinggi.

## Fitur Utama
- 📹 **Integrasi CCTV**: Mendukung input dari kamera IP maupun video lokal.
- 🧠 **Pendeteksian Anomali Otomatis**: Menggunakan model AI untuk mengenali aktivitas yang tidak biasa.
- ⚡ **Peringatan Real-Time**: Mengirim notifikasi atau alarm saat anomali terdeteksi.
- 💾 **Penyimpanan Bukti**: Menyimpan *snapshot* atau video klip anomali untuk analisis lebih lanjut.
- 🔍 **Visualisasi & Dashboard**: Tampilan antarmuka untuk memantau kamera dan log kejadian.

## Arsitektur Sistem
1. **Pengambilan Data Video** → dari CCTV / kamera IP.
2. **Pemrosesan Frame** → pra-pemrosesan gambar untuk model AI.
3. **Model AI** → klasifikasi dan deteksi anomali.
4. **Sistem Notifikasi** → mengirimkan peringatan ke admin.
5. **Penyimpanan Data** → log dan bukti video disimpan di server.

## Teknologi yang Digunakan
- Python 3.x
- OpenCV
- TensorFlow / PyTorch
- Flask / FastAPI (untuk API & dashboard)
- Message broker (contoh: RabbitMQ / Kafka) *(opsional)*
- Database (contoh: PostgreSQL / MongoDB)