function hule_mianzi(s, pai, i) {

    if (i == 9) return [[]];

    if (pai[i] == 0) return hule_mianzi(s, pai, i+1);

    var shunzi = [];
    if (i < 7 && pai[i] > 0 && pai[i+1] > 0 && pai[i+2] > 0) {
        pai[i]--; pai[i+1]--; pai[i+2]--;
        shunzi = hule_mianzi(s, pai, i);
        pai[i]++; pai[i+1]++; pai[i+2]++;
        for (var mianzi of shunzi) {
            mianzi.unshift(s+(i+1)+(i+2)+(i+3));
        }
    }

    var kezi = [];
    if (pai[i] >= 3) {
        pai[i] -= 3;
        kezi = hule_mianzi(s, pai, i);
        pai[i] += 3;
        for (var mianzi of kezi) {
            mianzi.unshift(s+(i+1)+(i+1)+(i+1));
        }
    }

    return shunzi.concat(kezi);
}

function hule_mianzi_all(shoupai) {

    var mianzi = [[]];

    for (var s of ['m','p','s']) {
        var new_mianzi = [];
        var sub_mianzi = hule_mianzi(s, shoupai._shouli[s], 0);
        for (var m of mianzi) {
            for (var n of sub_mianzi) {
                new_mianzi.push(m.concat(n));
            }
        }
        mianzi = new_mianzi;
    }

    var sub_mianzi_z = [];
    for (var n = 1; n <= 7; n++) {
        if (shoupai._shouli.z[n-1] == 0) continue;
        if (shoupai._shouli.z[n-1] != 3) return [];
        sub_mianzi_z.push('z'+n+n+n);
    }

    for (var i = 0; i < mianzi.length; i++) {
        mianzi[i] = mianzi[i].concat(sub_mianzi_z)
            .concat(shoupai._fulou);
    }

    return mianzi;
}

function add_hulepai(mianzi, hulepai) {

    var regexp   = new RegExp('^(' + hulepai[0] + '.*' + hulepai[1] +')');
    var replacer = '$1' + hulepai.substr(2) + '_';

    var add_mianzi = [];
    for (var i = 0; i < mianzi.length; i++) {
        if (mianzi[i].match(/[\-\+\=]/)) continue;
        if (i > 0 && mianzi[i] == mianzi[i-1]) continue;
        var rep = mianzi[i].replace(regexp, replacer);
        if (rep == mianzi[i]) continue;
        var new_mianzi = mianzi.concat();
        new_mianzi[i] = rep;
        add_mianzi.push(new_mianzi);
    }

    return add_mianzi;
}

function hule_yiban(shoupai, rongpai) {

    var hulepai = rongpai || shoupai._zimo;

    var hule_mianzi = [];
    for (var s in shoupai._shouli) {
        var pai = shoupai._shouli[s];
        for (var n = 1; n <= pai.length; n++) {
            if (pai[n-1] < 2) continue;
            var jiangpai = s+n+n;
            pai[n-1] -= 2;
            for (var mianzi of hule_mianzi_all(shoupai)) {
                mianzi.unshift(jiangpai);
                for (var add_mianzi of add_hulepai(mianzi, hulepai)) {
                    hule_mianzi.push(add_mianzi);
                }
            }
            pai[n-1] += 2;
        }
    }

    return hule_mianzi;
}

function hule_qiduizi(shoupai, rongpai) {

    if (shoupai._fulou.length > 0) return [];

    var hulepai = rongpai || shoupai._zimo;

    var hule_mianzi = [];
    for (var s in shoupai._shouli) {
        var pai = shoupai._shouli[s];
        for (var n = 1; n <= pai.length; n++) {
            if (pai[n-1] == 0) continue;
            if (pai[n-1] == 2) {
                var p = (s+n == hulepai.substr(0,2))
                    ? s+n+n + hulepai.substr(2) + '_'
                    : s+n+n;
                hule_mianzi.push(p);
            }
            else return [];
        }
    }

    return [hule_mianzi];
}

function hule_guoshi(shoupai, rongpai) {

    var hulepai = rongpai || shoupai._zimo;

    var hule_mianzi = [];
    for (var s in shoupai._shouli) {
        var pai = shoupai._shouli[s];
        var nn = s == 'z' ? [1,2,3,4,5,6,7] : [1,9];
        for (var n of nn) {
            if (pai[n-1] == 2) {
                var p = (s+n == hulepai.substr(0,2))
                    ? s+n+n + hulepai.substr(2) + '_'
                    : s+n+n;
                hule_mianzi.unshift(p);
            }
            else if (pai[n-1] == 1) {
                var p = (s+n == hulepai.substr(0,2))
                    ? s+n + hulepai.substr(2) + '_'
                    : s+n;
                hule_mianzi.push(p);
            }
            else return [];
        }
    }

    return [hule_mianzi];
}

function hule(shoupai, rongpai) {

    if (rongpai) {
        shoupai._zimo = rongpai.substr(0,2);
        shoupai._shouli[rongpai[0]][rongpai[1]-1]++;
    }

    return [].concat(hule_yiban(shoupai, rongpai))
        .concat(hule_qiduizi(shoupai, rongpai))
        .concat(hule_guoshi(shoupai, rongpai))
}