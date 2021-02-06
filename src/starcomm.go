package src

// Run is the starter function to start the App
func Run() {
	mapUrls()
	router.Run("0.0.0.0:8080")
}
