<?php
//
// 0.0.0.0 0.255.255.255   保留地址        保留地址
// 1.0.0.0 1.0.0.255       CLOUDFLARE.COM  CLOUDFLARE.COM          apnic.net
// 1.0.1.0 1.0.1.255       中国    福建                    电信
// 1.0.2.0 1.0.3.255       中国    福建                    电信
// 1.0.4.0 1.0.7.255       澳大利亚        维多利亚州      墨尔本          gtelecom.com.au
// 1.0.8.0 1.0.15.255      中国    广东                    电信
// 1.0.16.0        1.0.31.255      日本    日本                    i2ts.com
// 1.0.32.0        1.0.63.255      中国    广东                    电信
// 1.0.64.0        1.0.79.255      日本    广岛县                  megaegg.jp
// 1.0.80.0        1.0.95.255      日本    冈山县                  megaegg.jp
//
// 1.2.5.0	1.2.5.255	中国	北京	北京		电信
//

if(`ps uax |grep "iploc -port 8511"` == "") {
    die("start iploc first!\n");
}

$countries = include "country.php";
$zzcode = include "zzcode.php";
$areas = include "area.php";

$fp = fopen("ipip.txt", 'r');
$fpnew = fopen("newtext.txt", 'w');
$fpbad = fopen("newbad.txt", 'w');


$lastrow = null;

while($row = fgetcsv($fp, 0, "\t")) {
    $code = '';

    if(isset($areas[$row[2]])) {
        $code = $areas[$row[2]]['code'];
        if(isset($areas[$row[2]][$row[3]])) {
            $code = $areas[$row[2]][$row[3]]['code'];
            if(isset($areas[$row[2]][$row[3]])) {
                $code = $areas[$row[2]][$row[3]]['code'];
                if(isset($areas[$row[2]][$row[3]][$row[4]])) {
                    $code = $areas[$row[2]][$row[3]][$row[4]]['code'];
                }
            }
        }
    }
    elseif(isset($countries[$row[2]])){
        $code = $countries[$row[2]];
    }
    elseif(isset($zzcode[$row[2]])){
        $code = $zzcode[$row[2]];
    }
    else{
        $code = json_decode(file_get_contents("http://localhost:8511/iploc?ip={$row[0]}"),true)['loc'];
        fwrite($fpbad, sprintf("%s\t%s\t%s\n", $row[0], $row[1], $code));
    }

    if($lastrow==null){
        $lastrow = [$row[0], $row[1], $code];
    }
    else if($code==$lastrow[2]) {
        $lastrow[1] = $row[1];
    }
    else{
        fwrite($fpnew, join("\t",$lastrow));
        fwrite($fpnew, "\n");
        $lastrow = [$row[0], $row[1], $code];
    }
}

if($lastrow!=null){
    fwrite($fpnew, join("\t",$lastrow));
    fwrite($fpnew, "\n");
}

fclose($fpnew);
fclose($fp);

