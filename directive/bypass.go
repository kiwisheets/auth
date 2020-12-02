package directive

var bypass bool = false

func Development(isDevelopment bool) {
	bypass = isDevelopment
}
