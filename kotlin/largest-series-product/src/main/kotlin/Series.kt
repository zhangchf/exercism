/**
 * Created by zhangchf on 30/08/2017.
 */

class Series(val seriesString: String) {

    private val series by lazy(LazyThreadSafetyMode.SYNCHRONIZED) {
        seriesString.map { it.toString().toInt() }
    }

    init {
        require(seriesString.isNotBlank() && seriesString.all { it.isDigit() })
    }

    fun getLargestProduct(span: Int): Long {
        require(span > 0 && span <= series.size)

        return (0..(series.size-span)).map { product(it, span) }.max()!!
    }

    fun product(from: Int, span: Int): Long =
            series.slice(from until from + span).fold(1L, {acc, v -> acc * v})

}