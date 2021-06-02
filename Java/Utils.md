# 常用方法

## 1、对消耗时间友好的处理

> 毫秒 --> 友好的时间  xh xm xs 


```xml
        <dependency>
            <groupId>joda-time</groupId>
            <artifactId>joda-time</artifactId>
            <version>2.10.10</version>
        </dependency>
```

```java
public class TimerFormat {

    public static String TimerFormats(Long timer){
        Duration duration = new Duration(timer); // in milliseconds
        PeriodFormatter formatter = new PeriodFormatterBuilder()
                .appendDays()
                .appendSuffix("d ")
                .appendHours()
                .appendSuffix("h ")
                .appendMinutes()
                .appendSuffix("m ")
                .appendSeconds()
                .appendSuffix("s")
                .toFormatter();
        String formatted = formatter.print(duration.toPeriod());
        return formatted;
    }

    
}
```

## 2、字节单位转换

```java

public class ConvertSize {
    public  static  String Convert(long size) {
        //如果字节数少于1024，则直接以B为单位，否则先除于1024，后3位因太少无意义
        if (size < 1024) {
            return String.valueOf(size) + "B";
        } else {
            size = size / 1024;
        }
        //如果原字节数除于1024之后，少于1024，则可以直接以KB作为单位
        //因为还没有到达要使用另一个单位的时候
        //接下去以此类推
        if (size < 1024) {
            return String.valueOf(size) + "KB";
        } else {
            size = size / 1024;
        }
        if (size < 1024) {
            //因为如果以MB为单位的话，要保留最后1位小数，
            //因此，把此数乘以100之后再取余
            size = size * 100;
            return String.valueOf((size / 100)) + "."
                    + String.valueOf((size % 100)) + "MB";
        } else {
            //否则如果要以GB为单位的，先除于1024再作同样的处理
            size = size * 100 / 1024;
            return String.valueOf((size / 100)) + "."
                    + String.valueOf((size % 100)) + "GB";
        }
    }
}

```



## 3、将列表按照某个属性分组

```xml
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>fastjson</artifactId>
            <version>1.2.73</version>
        </dependency>
```

```java
// entity

@Data
public class Coupon {
    private Integer couponId;
    private Integer price;
    private String name;
}

// service
public class ListGroupTest {
    public static void main(String[] args) {
        List<Coupon> couponList = new ArrayList<>();
        Coupon coupon1 = new Coupon(1,100,"优惠券1");
        Coupon coupon2 = new Coupon(2,200,"优惠券2");
        Coupon coupon3 = new Coupon(3,300,"优惠券3");
        Coupon coupon4 = new Coupon(3,400,"优惠券4");
        couponList.add(coupon1);
        couponList.add(coupon2);
        couponList.add(coupon3);
        couponList.add(coupon4);

        Map<Integer, List<Coupon>> resultList = couponList.stream().collect(Collectors.groupingBy(Coupon::getCouponId));
        System.out.println(JSON.toJSONString(resultList, SerializerFeature.PrettyFormat));
    }
}

// out
{
	1:[
			{
				"couponId":1,
				"name":"优惠券1",
				"price":100
			}
	  ],
	2:[
			{
				"couponId":2,
				"name":"优惠券2",
				"price":200
			}
	  ],
	3:[
			{
				"couponId":3,
				"name":"优惠券3",
				"price":300
			},
			{
				"couponId":3,
				"name":"优惠券4",
				"price":400
			}
	  ]
}


// 如果分组后，分组内并不想是对象，而是对象的属性，也可以做到的。
public class ListGroupTest2 {
    public static void main(String[] args) {
        List<Coupon> couponList = new ArrayList<>();
        Coupon coupon1 = new Coupon(1,100,"优惠券1");
        Coupon coupon2 = new Coupon(2,200,"优惠券2");
        Coupon coupon3 = new Coupon(3,300,"优惠券3");
        Coupon coupon4 = new Coupon(3,400,"优惠券4");
        couponList.add(coupon1);
        couponList.add(coupon2);
        couponList.add(coupon3);
        couponList.add(coupon4);

        Map<Integer, List<String>> resultList = couponList.stream().collect(Collectors.groupingBy(Coupon::getCouponId,Collectors.mapping(Coupon::getName,Collectors.toList())));
        System.out.println(JSON.toJSONString(resultList, SerializerFeature.PrettyFormat));
    }
}
 
// out
{
	1:[
		"优惠券1"
	  ],
	2:[
		"优惠券2"
	  ],
	3:[
		"优惠券3",
		"优惠券4"
	  ]
}


```

