
enum class Classification {
    DEFICIENT, PERFECT, ABUNDANT
}

fun classify(naturalNumber: Int): Classification {

    require(naturalNumber > 0)

    val aliquotSum = (1 .. naturalNumber / 2).filter { naturalNumber.rem(it) == 0 }.sum()
    val difference = aliquotSum - naturalNumber
    return when {
        difference > 0 -> Classification.ABUNDANT
        difference == 0 -> Classification.PERFECT
        else -> Classification.DEFICIENT
    }
}
