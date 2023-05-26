const puppeteer = require('puppeteer');

async function inner(url) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    page.on('console', msg => {
        // console.log(msg._text);
    });
    const device = {
        userAgent: 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36',
        viewport: {
            width: 1366,
            height: 768
        }
    };
    await page.emulate(device);
    await page.goto(url);
    const texts = await page.evaluate(() => {
        const END_WITH_PUNCTUATION_REGEX = /[\.\,\?\!\'\"\:\;\-\—\。\？\！\，\、\；\：\“\”\﹃\﹄\﹁\﹂\（\）\［\］\〔\〕\【\】\—\…\《\》\〈\〉\﹏\＿]$/;

        function FullTexttoString(array) {
            let string = '';
            let sentence = '';
            let temp = [];
            let sent = '';
            let i = 0;
            // console.log(JSON.stringify(array));

            array.forEach(function(obj, indexArray) {
                if (!obj.trEnd && !obj.preEnd) {
                    if (obj.inPre) {
                        sentence = obj.text.replace(/[\t\f\ ]+/g, ' ');
                        sentence = sentence.trim();
                        if (sentence.length) {
                            temp = sentence.split('\n').filter(str => str.trim().length);

                            temp.forEach(function(sent, indexTemp) {
                                sent = sent.trim();
                                if (sent.length) {
                                    if (obj.inH) {
                                        if (!sent.endsWith('-')) {
                                            string += sent + ' -' + '\n';
                                        } else {
                                            string += sent + '\n';
                                        }
                                    } else if (!obj.inTr && !obj.inTdTh) {
                                        if (indexTemp === temp.length - 1 && indexArray < array.length - 1) {
                                            i = indexArray + 1;
                                            while (array[i].trEnd) {
                                                i++;
                                            }
                                            if (array[i].preEnd) {
                                                if (!END_WITH_PUNCTUATION_REGEX.test(sent)) {
                                                    string += sent + '.' + '\n';
                                                } else {
                                                    string += sent + '\n';
                                                }
                                            } else {
                                                string += sent + '\n';
                                            }
                                        } else {
                                            string += sent + '\n';
                                        }
                                    } else {
                                        if (indexTemp === temp.length - 1 && indexArray < array.length - 1) {
                                            i = indexArray + 1;
                                            while (array[i].preEnd) {
                                                i++;
                                            }
                                            if (array[i].trEnd) {
                                                if (!END_WITH_PUNCTUATION_REGEX.test(sent)) {
                                                    string += sent + '.' + '\n';
                                                } else {
                                                    string += sent + '\n';
                                                }
                                            } else {
                                                string += sent + '\n';
                                            }
                                        } else {
                                            string += sent + '\n';
                                        }
                                    }
                                }
                            });
                        }
                    } else {
                        sentence = obj.text.replace(/\s+/g, ' ');
                        sentence = sentence.trim();
                        if (sentence.length) {
                            if (obj.inH) {
                                if (!sent.endsWith('-')) {
                                    string += sentence + ' -' + '\n';
                                } else {
                                    string += sentence + '\n';
                                }
                            } else if (!obj.inTdTh && !obj.inTr) {
                                if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                                    string += sentence + '.' + '\n';
                                } else {
                                    string += sentence + '\n';
                                }
                            } else {
                                if (indexArray < array.length - 1) {
                                    if (array[indexArray + 1].trEnd) {
                                        if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                                            string += sentence + '.' + '\n';
                                        } else {
                                            string += sentence + '\n';
                                        }
                                    } else {
                                        if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                                            string += sentence + ',' + '\n';
                                        } else {
                                            string += sentence + '\n';
                                        }
                                    }
                                } else {
                                    if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                                        string += sentence + ',' + '\n';
                                    } else {
                                        string += sentence + '\n';
                                    }
                                }
                            }
                        }
                    }
                }
            });
            return string;
        }

        function getDomNodeFullText(dom) {
            const body_stack = [null];

            let text = [];

            let sentence = "";

            let childrens = [];

            let displayType = 'block';

            let invisible = false;

            let positionType = '';

            let floatType = '';

            let inPre = false;

            let inTdTh = false;

            let inTr = false;

            let inH = false;

            let i = 0;

            let temp = '';

            const INLINE_ELEMS = [
                'SPAN', 'A', 'EM', 'STRONG', 'CODE', 'SAMP',
                'KBD', 'VAR', 'I', 'U', 'B', 'INS', 'TT',
                'ABBR', 'ACRONYM', 'BDI', 'BDO', 'BIG', 'CITE',
                'DFN', 'INS', 'LABEL', 'MARK', 'SMALL', 'SUB',
                'SUP', 'TIME', 'WBR'
            ];

            const BLOCK_ELEMS = [
                'DIV', 'P', 'H1', 'H2', 'H3', 'H4', 'H5',
                'H6', 'BLOCKQUOTE', 'DL', 'DT', 'DD',
                'THEAD', 'TFOOT', 'TBODY', 'TR', 'UL', 'TABLE',
                'OL', 'LI', 'ADDRESS', 'ARTICLE', 'ASIDE',
                'CAPTION', 'CENTER', 'DETAILS', 'SUMMARY',
                'DIR', 'FORM', 'LEGEND', 'FIELDSET', 'FIGURE',
                'FIGCAPTION', 'FOOTER', 'HEADER', 'HGROUP', 'MAIN',
                'NAV', 'SECTION', 'TEXTAREA', 'TD', 'TH', 'PRE'
            ];

            const TABLE_TD_TH_ELEMS = [
                'TH', 'TD'
            ];

            const HEADER_ELEMS = [
                'H1', 'H2', 'H3', 'H4', 'H5', 'H6'
            ];

            const BREAK_ELEMS = [
                'BR', 'HR'
            ];

            childrens = dom.childNodes;

            childrens.forEach(function(node) {
                body_stack.push(node);
            });

            while (body_stack.length) {
                let node = body_stack.pop();
                if (node && node.nodeType && node.tagName && node.tagName === 'TR') {
                    console.log('hit');
                }
                if (node === null) { // Generate sentence
                    // console.log(JSON.stringify(sentence));
                    temp = sentence.trim();
                    if (temp) {
                        text.unshift({
                            inPre: inPre,
                            inTdTh: inTdTh,
                            inTr: inTr,
                            inH: inH,
                            trEnd: false,
                            preEnd: false,
                            text: temp
                        });
                    }

                    // if (inPre) {
                    //     sentence = sentence.trim();
                    //     if (sentence.length) {
                    //         if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                    //             sentence = sentence + '.';
                    //         }
                    //         temp = sentence.split('\n');
                    //         temp.reverse().forEach(function(sent) {
                    //             sent = sent.trim();
                    //             if (sent.length) {
                    //                 text.unshift(sent);
                    //             }
                    //         });
                    //     }
                    // } else {
                    //     sentence = sentence.replace(/\s+/g, ' ');
                    //     sentence = sentence.trim();
                    //     if (sentence.length) {
                    //         if (!END_WITH_PUNCTUATION_REGEX.test(sentence)) {
                    //             sentence = sentence + '.';
                    //         }
                    //         text.unshift(sentence);
                    //     }
                    // }
                    sentence = "";
                } else if (node === 'pre start') { // Pre Element Start
                    inPre = true;
                    text.unshift({
                        preEnd: true
                    });
                } else if (node === 'pre end') { // Pre Element End
                    inPre = false;
                } else if (node === 'tr start') { // Tr Element Start
                    inTr = true;
                    text.unshift({
                        trEnd: true
                    });
                } else if (node === 'tr end') { // Tr Element End
                    inTr = false;
                } else if (node === 'td/th start') { // Td Th Element Start
                    inTdTh = true;
                } else if (node === 'td/th end') { // Td Th Element End
                    inTdTh = false;
                } else if (node === 'h start') { // H Element Start
                    inH = true;
                } else if (node === 'h end') { // H Element End
                    inH = false;
                } else if (node.nodeType === 3) { // Text Node
                    sentence = node.textContent + sentence;
                } else if (node.nodeType === 1) { // Element Node
                    displayType = window.getComputedStyle(node, null).getPropertyValue('display').trim().toLowerCase();
                    invisible = window.getComputedStyle(node, null).getPropertyValue('opacity') === "0" || window.getComputedStyle(node, null).getPropertyValue('visibility') === "hidden";
                    positionType = window.getComputedStyle(node, null).getPropertyValue('position').trim().toLowerCase();
                    floatType = window.getComputedStyle(node, null).getPropertyValue('float').trim().toLowerCase();
                    if (displayType != 'none' && !invisible) {
                        if (BLOCK_ELEMS.indexOf(node.tagName) > -1 || displayType === 'block' || displayType === 'flex' || positionType === 'absolute' || floatType === 'left' || floatType === 'right') { // Block Node
                            if (node.tagName === 'PRE') { // Pre Node
                                if (!inPre) {
                                    body_stack.push('pre end');
                                }
                                body_stack.push(null);
                                childrens = node.childNodes;
                                childrens.forEach(function(node) {
                                    body_stack.push(node);
                                });
                                body_stack.push('pre start');
                                body_stack.push(null);
                            } else if (node.tagName === 'TR') {
                                // console.log(JSON.stringify(node.innerText));
                                body_stack.push('tr end');
                                body_stack.push(null);
                                childrens = node.childNodes;
                                childrens.forEach(function(node) {
                                    body_stack.push(node);
                                });
                                body_stack.push('tr start');
                                body_stack.push(null);
                            } else if (TABLE_TD_TH_ELEMS.indexOf(node.tagName) > -1) {
                                body_stack.push('td/th end');
                                body_stack.push(null);
                                childrens = node.childNodes;
                                childrens.forEach(function(node) {
                                    body_stack.push(node);
                                });
                                body_stack.push('td/th start');
                                body_stack.push(null);
                            } else if (HEADER_ELEMS.indexOf(node.tagName) > -1) {
                                body_stack.push('h end');
                                body_stack.push(null);
                                childrens = node.childNodes;
                                childrens.forEach(function(node) {
                                    body_stack.push(node);
                                });
                                body_stack.push('h start');
                                body_stack.push(null);
                            } else { // Not Special Node
                                body_stack.push(null);
                                childrens = node.childNodes;
                                childrens.forEach(function(node) {
                                    body_stack.push(node);
                                });
                                body_stack.push(null);
                            }
                        } else if (INLINE_ELEMS.indexOf(node.tagName) > -1) { // Inline Node
                            childrens = node.childNodes;
                            childrens.forEach(function(node) {
                                body_stack.push(node);
                            });
                        } else if (BREAK_ELEMS.indexOf(node.tagName) > -1) { // Breakline Node
                            body_stack.push(null);
                        } else if (node.tagName === 'Q') { // Quote Node
                            body_stack.push(window.document.createTextNode('"'));
                            childrens = node.childNodes;
                            childrens.forEach(function(node) {
                                body_stack.push(node);
                            });
                            body_stack.push(window.document.createTextNode('"'));
                        }
                    }
                }
            }

            return FullTexttoString(text);
        }

        return {
            og: window.document.querySelector("meta[property='og:title']") && window.document.querySelector("meta[property='og:title']").getAttribute('content'),
            title: window.document.title,
            desc: window.document.querySelector('meta[name="description"]') && window.document.querySelector('meta[name="description"]').getAttribute('content') || window.document.querySelector("meta[property='og:description']") && window.document.querySelector("meta[property='og:description']").getAttribute('content'),
            fulltext: getDomNodeFullText(window.document.body)
        };
    });

    // console.log(texts.fulltext);

    await browser.close();
    // let r, strBody, charset, tmpCharset, bodyDecoded;
    // r = await to(request({
    //     url: "http://127.0.0.1:8001/api/spider/collect/get",
    //     headers: {
    //         'Secret-Key': 'amVmZmdlZWs='
    //     },
    //     timeout: 5000,
    //     encoding: null,
    //     strictSSL: false,
    //     json: {"user_id": user_id, "url_hash": url_hash}
    // }));
    // if (r[1]) {
    //     // http get error
    //     return [null, {errcode: 40403, errmsg: '请求失败'}];
    // } else {
    //     if (r[0].statusCode !== 200) {
    //         return [null, {errcode: 40405, errmsg: '状态码错误'}];
    //     }
    //     // http success
    //     return [r[0].body, null];
    // }
    return texts;
};

exports.get = async function(url) {
    try {
        var texts = await inner(url);
    } catch (error) {
        return [null, error];
    }  
    return [texts, null];
};