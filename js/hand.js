function hupai(mianzi, zhuangfeng, zifeng) {

    var menqian = mianzi.filter(function(m){
        return m.match(/[\-\+\=](?!_)/)}).length == 0;
    // 門前のとき true

    var zhuangfengpai   = new RegExp('^z' + ((zhuangfeng + 1) % 4) + '.*$');
    // 場風
    var zifengpai       = new RegExp('^z' + ((zifeng + 1) % 4) + '.*$');
    // 自風
    var fengpai         = /^z[1234].*$/;                // 風牌
    var sanyuanpai      = /^z[567].*$/;                 // 三元牌

    var yaojiu          = /^.*[z19].*$/;                // 幺九牌
    var zipai           = /^z.*$/;                      // 字牌

    var shunzi          = /^[mpsz](?!(\d)\1).*$/;       // 順子
    var kezi            = /^[mpsz](\d)\1\1.*$/;         // 刻子
    var ankezi          = /^[mpsz](\d)\1\1(?:\1|_)?$/;  // 暗刻子
    var gangzi          = /^[mpsz](\d)\1\1.*\1.*$/;     // 槓子

    var danqi           = /^[mpsz](\d)\1[\-\+\=]?_$/;   // 単騎待ち
    var kanzhang        = /^[mps]\d\d[\-\+\=]?_\d$/;    // 嵌張待ち
    var bianzhang       = /^[mps](123[\-\+\=]?_|7[\-\+\=]?_89)$/;
    // 辺張待ち

    function menqianqing() {
        if (mianzi.filter(function(m){return m.match(/[\-\+\=]/)}).length > 0)
            return [];
        return [{ name: '門前清自摸和', fanshu: 1 }];
    }
    function fanpai() {
        if (mianzi.length != 5)             return [];
        var hupai_all = [];
        if (mianzi.filter(function(m){
            return m.match(kezi) && m.match(zhuangfengpai)
        }).length > 0)
            hupai_all.push({ name: '翻牌 場風', fanshu: 1 });
        if (mianzi.filter(function(m){
            return m.match(kezi) && m.match(zifengpai)
        }).length > 0)
            hupai_all.push({ name: '翻牌 自風', fanshu: 1 });
        if (mianzi.filter(function(m){
            return m.match(kezi) && m.match(/^z5.*$/)
        }).length > 0)
            hupai_all.push({ name: '翻牌 白', fanshu: 1 });
        if (mianzi.filter(function(m){
            return m.match(kezi) && m.match(/^z6.*$/)
        }).length > 0)
            hupai_all.push({ name: '翻牌 發', fanshu: 1 });
        if (mianzi.filter(function(m){
            return m.match(kezi) && m.match(/^z7.*$/)
        }).length > 0)
            hupai_all.push({ name: '翻牌 中', fanshu: 1 });
        return hupai_all;
    }
    function pinghu() {
        if (mianzi.length != 5)             return [];
        if (! menqian)                      return [];
        if (mianzi[0].match(zhuangfengpai)) return [];
        if (mianzi[0].match(zifengpai))     return [];
        if (mianzi[0].match(sanyuanpai))    return [];
        if (mianzi.filter(function(m){return m.match(kezi)}).length > 0)
            return [];
        if (mianzi.filter(function(m){return m.match(danqi)}).length > 0)
            return [];
        if (mianzi.filter(function(m){return m.match(kanzhang)}).length > 0)
            return [];
        if (mianzi.filter(function(m){return m.match(bianzhang)}).length > 0)
            return [];
        return [{ name: '平和', fanshu: 1 }];
    }
    function duanyaojiu() {
        if (mianzi.filter(function(m){return m.match(yaojiu)}).length > 0)
            return [];
        return [{ name: '断幺九', fanshu: 1 }];
    }
    function yibeikou() {
        if (mianzi.length != 5)             return [];
        if (! menqian)                      return [];
        var map = {};
        for (var mm of mianzi.filter(function(m){return m.match(shunzi)})) {
            mm = mm.replace(/[\-\+\=]?_/, '');
            if (! map[mm]) map[mm] = 1;
            else           map[mm]++;
        }
        var beikou = 0;
        for (var mm in map) {
            if (map[mm] > 3) beikou++;
            if (map[mm] > 1) beikou++;
        }
        if (beikou != 1)                    return [];
        return [{ name: '一盃口', fanshu: 1 }];
    }
    function sansetongshun() {
        if (mianzi.length != 5)             return [];
        var map = { m: {}, p: {}, s: {} }
        for (var mm of mianzi.filter(function(m){return m.match(shunzi)})) {
            mm = mm.replace(/[\-\+\=\_]/g, '');
            map[mm[0]][mm.substr(1)] = 1;
        }
        for (var mm in map.m) {
            if (map.p[mm] && map.s[mm])
                return [{ name: '三色同順', fanshu: (menqian ? 2 : 1) }];
        }
        return [];
    }
    function yiqitongguan() {
        if (mianzi.length != 5)             return [];
        var map = { m: {}, p: {}, s: {} }
        for (var mm of mianzi.filter(function(m){return m.match(shunzi)})) {
            mm = mm.replace(/[\-\+\=\_]/g, '');
            map[mm[0]][mm.substr(1)] = 1;
        }
        for (var s in map) {
            if (map[s][123] && map[s][456] && map[s][789])
                return [{ name: '一気通貫', fanshu: (menqian ? 2 : 1) }];
        }
        return [];
    }
    function hunquandaiyaojiu() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(yaojiu)}).length != 5)
            return [];
        if (mianzi.filter(function(m){return m.match(shunzi)}).length == 0)
            return [];
        if (mianzi.filter(function(m){return m.match(zipai)}).length == 0)
            return [];
        return [{ name: '混全帯幺九', fanshu: (menqian ? 2 : 1) }];
    }
    function qiduizi() {
        if (mianzi.length != 7)             return [];
        return [{ name: '七対子', fanshu: 2 }];
    }
    function duiduihu() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(kezi)}).length != 4)
            return [];
        return [{ name: '対々和', fanshu: 2 }];
    }
    function sananke() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(ankezi)}).length != 3)
            return [];
        return [{ name: '三暗刻', fanshu: 2 }];
    }
    function sangangzi() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(gangzi)}).length != 3)
            return [];
        return [{ name: '三槓子', fanshu: 2 }];
    }
    function sansetongke() {
        if (mianzi.length != 5)             return [];
        var map = { m: {}, p: {}, s: {}, z: {} }
        for (var mm of mianzi.filter(function(m){return m.match(kezi)})) {
            map[mm[0]][mm.substr(1,3)] = 1;
        }
        for (var mm in map.m) {
            if (map.p[mm] && map.s[mm])
                return [{ name: '三色同刻', fanshu: 2 }];
        }
        return [];
    }
    function hunlaotou() {
        if (mianzi.filter(function(m){return ! m.match(yaojiu)}).length > 0)
            return [];
        if (mianzi.filter(function(m){return m.match(shunzi)}).length > 0)
            return [];
        if (mianzi.filter(function(m){return m.match(zipai)}).length == 0)
            return [];
        return [{ name: '混老頭', fanshu: 2 }];
    }
    function xiaosanyuan() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(sanyuanpai)}).length != 3)
            return [];
        if (! mianzi[0].match(sanyuanpai))  return [];
        return [{ name: '小三元', fanshu: 2 }];
    }
    function hunyise() {
        if (mianzi.filter(function(m){return m.match(zipai)}).length == 0)
            return [];
        for (var s of ['m','p','s']) {
            var yise = new RegExp('^[z' + s + '].*$')
            if (mianzi.filter(function(m){return m.match(yise)}).length
                == mianzi.length)
                return [{ name: '混一色', fanshu: (menqian ? 3 : 2) }];
        }
        return [];
    }
    function chunquandaiyaojiu() {
        if (mianzi.length != 5)             return [];
        if (mianzi.filter(function(m){return m.match(yaojiu)}).length != 5)
            return [];
        if (mianzi.filter(function(m){return m.match(shunzi)}).length == 0)
            return [];
        if (mianzi.filter(function(m){return m.match(zipai)}).length > 0)
            return [];
        return [{ name: '純全帯幺九', fanshu: (menqian ? 3 : 2) }];
    }
    function erbeikou() {
        if (mianzi.length != 5)             return [];
        if (! menqian)                      return [];
        var map = {};
        for (var mm of mianzi.filter(function(m){return m.match(shunzi)})) {
            mm = mm.replace(/[\-\+\=]?_/, '');
            if (! map[mm]) map[mm] = 1;
            else           map[mm]++;
        }
        var beikou = 0;
        for (var mm in map) {
            if (map[mm] > 3) beikou++;
            if (map[mm] > 1) beikou++;
        }
        if (beikou != 2)                    return [];
        return [{ name: '二盃口', fanshu: 3 }];
    }
    function qingyise() {
        for (var s of ['m','p','s']) {
            var yise = new RegExp('^' + s + '.*$')
            if (mianzi.filter(function(m){return m.match(yise)}).length
                == mianzi.length)
                return [{ name: '清一色', fanshu: (menqian ? 6 : 5) }];
        }
        return [];
    }

    return    [].concat(menqianqing())
        .concat(fanpai())
        .concat(pinghu())
        .concat(duanyaojiu())
        .concat(yibeikou())
        .concat(sansetongshun())
        .concat(yiqitongguan())
        .concat(hunquandaiyaojiu())
        .concat(qiduizi())
        .concat(duiduihu())
        .concat(sananke())
        .concat(sangangzi())
        .concat(sansetongke())
        .concat(hunlaotou())
        .concat(xiaosanyuan())
        .concat(hunyise())
        .concat(chunquandaiyaojiu())
        .concat(erbeikou())
        .concat(qingyise())
}