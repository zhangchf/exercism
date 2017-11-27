/**
 * Created by zhangchf on 29/08/2017.
 */
object Raindrops {
/*    fun convert(input: Int): String {
        var result = ""
        if (input.hasFactor(3)) result += "Pling"
        if (input.hasFactor(5)) result += "Plang"
        if (input.hasFactor(7)) result += "Plong"
        if (result.isBlank()) result = input.toString()
        return result
    }*/

    fun convert(input: Int): String = buildString {

        if (input.hasFactor(3)) append("Pling")
        if (input.hasFactor(5)) append("Plang")
        if (input.hasFactor(7)) append("Plong")
        if (isBlank()) append(input.toString())
    }

    fun Int.hasFactor(factor: Int): Boolean = this.rem(factor) == 0
}