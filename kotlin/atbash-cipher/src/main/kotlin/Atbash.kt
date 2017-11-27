/**
 * Created by zhangchf on 31/08/2017.
 */

object Atbash {

    val keyMapping = ('a'..'z').zip('z' downTo 'a').toMap()

    fun encode(input: String): String {
        return input.filter { it.isLetterOrDigit() }
                .toLowerCase()
                .map { if (it.isLetter()) keyMapping[it] else it }
                .batch(5)
                .map { it.joinToString(separator = "") }
                .joinToString(separator = " ").trim()
    }

    fun decode(input: String): String {
        return input.filter { it.isLetterOrDigit() }
                .toLowerCase()
                .map { if (it.isLetter()) keyMapping[it] else it }.joinToString(separator = "")
    }

    private fun <E> List<E>.batch(batchSize: Int): List<List<E>> =
        when {
            this.size <= batchSize -> listOf(this)
            else -> listOf(this.take(batchSize)) + this.drop(batchSize).batch(batchSize)
        }
}