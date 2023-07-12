package seeds

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

func SeedTodos(db *gorm.DB) error {
	todos := []models.Todo{
		{Name: "Melakukan check-up rutin selama masa kehamilan"},
		{Name: "Memastikan asupan nutrisi yang cukup selama masa kehamilan"},
		{Name: "Melakukan imunisasi lengkap setelah kelahiran"},
		{Name: "Menyusui eksklusif hingga usia 6 bulan"},
		{Name: "Memperkenalkan makanan pendamping ASI setelah 6 bulan"},
		{Name: "Memastikan asupan nutrisi seimbang pada anak hingga usia 9 tahun"},
		{Name: "Melakukan imunisasi lanjutan sesuai jadwal"},
		{Name: "Rutin melakukan check-up kesehatan anak"},
		{Name: "Memastikan perkembangan fisik dan mental anak sesuai usia"},
		{Name: "Mengajarkan kebiasaan hidup sehat pada anak"},
		{Name: "Konsumsi vitamin prenatal harian"},
		{Name: "Konsumsi makanan seimbang dan bergizi untuk ibu hamil"},
		{Name: "Melakukan olahraga ringan seperti berjalan selama 30 menit"},
		{Name: "Berikan ASI eksklusif bagi bayi usia dibawah 6 bulan"},
		{Name: "Periksa berat badan dan panjang badan anak secara rutin"},
		{Name: "Cek perkembangan kognitif dan motorik anak"},
		{Name: "Pastikan anak tidur dengan cukup"},
		{Name: "Berikan MP-ASI setelah anak berusia 6 bulan"},
		{Name: "Berikan anak vaksin sesuai jadwal imunisasi"},
		{Name: "Mendongeng sebelum anak tidur"},
		{Name: "Berikan perhatian dan cinta pada anak"},
		{Name: "Perkenalkan buah dan sayur pada anak usia 7 bulan keatas"},
		{Name: "Bersihkan mainan anak secara rutin"},
		{Name: "Ajarkan anak untuk mencuci tangan dengan benar"},
		{Name: "Lakukan aktivitas fisik bersama anak"},
		{Name: "Perkenalkan makanan baru pada anak"},
		{Name: "Baca buku bersama anak"},
		{Name: "Melakukan pemeriksaan rutin ke dokter anak"},
		{Name: "Bermain dan berinteraksi dengan anak"},
		{Name: "Ajarkan anak untuk makan dan minum sendiri"},
		{Name: "Pastikan anak cukup minum air putih setiap hari"},
		{Name: "Ajarkan anak kebersihan diri"},
		{Name: "Berikan pujian positif terhadap perkembangan anak"},
		{Name: "Ajarkan anak untuk berbagi dan bersosialisasi"},
		{Name: "Membantu anak dalam belajar berbicara"},
		{Name: "Ajarkan anak mengenai nama-nama bagian tubuh"},
		{Name: "Ajarkan anak warna dan bentuk"},
		{Name: "Pergi ke taman bermain bersama anak"},
		{Name: "Melakukan aktivitas seni dan kreatifitas bersama anak"},
		{Name: "Pantau jadwal tidur anak"},
		{Name: "Bantu anak belajar mengunyah makanan"},
		{Name: "Ajarkan anak untuk berterima kasih"},
		{Name: "Bacakan cerita atau dongeng sebelum tidur"},
		{Name: "Bantu anak belajar mengenali emosinya"},
		{Name: "Bantu anak belajar mengekspresikan dirinya"},
		{Name: "Melakukan aktivitas menyanyi dan menari bersama anak"},
		{Name: "Berikan anak waktu luang untuk bermain sendiri"},
		{Name: "Pantau tayangan yang ditonton anak"},
		{Name: "Pergi berbelanja kebutuhan anak bersama anak"},
		{Name: "Ajarkan anak untuk mengucapkan permisi"},
		{Name: "Ajarkan anak untuk menunggu giliran"},
		{Name: "Bantu anak belajar mengenali angka dan abjad"},
		{Name: "Lakukan permainan edukatif bersama anak"},
		{Name: "Bantu anak belajar berpakaian sendiri"},
		{Name: "Ajarkan anak untuk merapikan mainannya sendiri"},
		{Name: "Ajarkan anak untuk menjaga lingkungan"},
		{Name: "Ajarkan anak untuk mengepak barang-barangnya sendiri"},
		{Name: "Melakukan perjalanan pendek atau piknik bersama anak"},
		{Name: "Bantu anak belajar memasang sabuk pengaman di mobil"},
		{Name: "Ajarkan anak untuk makan dengan sendok dan garpu"},
		{Name: "Ajarkan anak untuk menggosok gigi secara mandiri"},
	}

	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			return err
		}
	}

	return nil
}
