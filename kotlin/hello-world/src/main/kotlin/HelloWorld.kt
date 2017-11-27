fun hello(name: String = ""): String {
   return "Hello, ${whom(name)}!"
}

private fun whom(name: String): String {
    return if (name.isBlank()) "World" else name
}