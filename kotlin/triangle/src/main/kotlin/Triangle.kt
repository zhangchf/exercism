/**
 * Created by zhangchf on 29/08/2017.
 */

open class Polygon(vararg val edges: Number) {
    init {
        require(edges.size >= 3)
        require(edges.all { it.toDouble() > 0 })
    }

    val differentLengths = edges.distinct().count()
    fun edge(index: Int) = edges[index % edges.size].toDouble()
}

class Triangle(a: Number, b: Number, c: Number): Polygon(a, b, c) {
    init {
        require(
                (0..2).all { edge(it) <= edge(it+1) + edge(it+2) }
        )
    }

    val isEquilateral
        get() = differentLengths == 1

    val isIsosceles = differentLengths < 3
    val isScalene = differentLengths == 3
    val isDegenerate = (2 * edges.maxBy { it.toDouble() }!!.toDouble() == edges.sumByDouble { it.toDouble() })
}


/*

class Triangle(val a: Number, val b: Double, val c: Double) {

    constructor(a: Int, b: Int, c: Int): this(a.toDouble(), b.toDouble(), c.toDouble())

    init {
        require(isValid())
    }

    val isEquilateral = isValid() && (a == b && a == c)
    val isIsosceles = isValid() && (a == b || b == c || a == c)
    val isScalene = isValid() && (a != b && a != c && b != c)

    private fun isValid(): Boolean {
        if (a <= 0 || b <= 0 || c <= 0) return false
        if (a + b < c || a + c < b || b + c < a) return false
        return true
    }
}*/
