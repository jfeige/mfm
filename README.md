# mfm

之前在golang中操作mysql返回的结果集时，都是调用Scan(dst...)方法读取字段值，如果字段比较多的话，就很繁琐<br/>但是golang的mysql驱动本身不提供获取结果集的数值，只提供返回结果集的字段名列表，后来偶尔在网上看到某种解析的方法，当时忙着赶进度，只看了个大概，趁着今天有时间，赶紧写下来。
