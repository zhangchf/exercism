/**
 * Created by zhangchf on 04/09/2017.
 */

class Matrix(val points: List<List<Int>>) {
    val rows: Int
    val cols: Int

    init {
        require(points.size > 0 && (0 until points.size).all { points[it].size == points[0].size }) {"Invalide init matrix points list"}
        rows = points.size
        cols = points[0].size
    }

    var saddlePoints: Set<MatrixCoordinate>? = null
    get() {
        var result = mutableSetOf<MatrixCoordinate>()

        (0 until rows).map { row ->
            (0 until cols).filter { col ->
                val item = points[row][col]
                points[row].all { it <= item } && (0 until rows).all { points[it][col] >= item }
            }.forEach { result.add(MatrixCoordinate(row, it)) }
        }
        return result
    }

}