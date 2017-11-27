import java.time.Duration
import java.time.LocalDate
import java.time.LocalDateTime
import java.time.LocalTime

/**
 * Created by zhangchf on 28/08/2017.
 */
class Gigasecond(localDateTime: LocalDateTime) {

    constructor(localDate: LocalDate): this(localDate.atStartOfDay())

    private val GIGASECONDS: Long = Math.pow(10.0, 9.0).toLong()

    var date: LocalDateTime = localDateTime.plusSeconds(GIGASECONDS)

}