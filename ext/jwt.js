const crypto = require('crypto')

const convertBase64 = json => {
    const jsonStr = JSON.stringify(json)
    const jsonB64 = Buffer.from(jsonStr).toString('base64')
    const jsonB64NoPadding = jsonB64.replace(/={1,2}$/, '')
    console.log('Json: ' + jsonStr);
    console.log('Base64: ' + jsonB64);
    console.log('Base64(No Padding): ' + jsonB64NoPadding);
    return jsonB64NoPadding
}

const encryptHmacSHA256 = (secretKey, unsignedToken) => {
    const hash = crypto.createHmac('sha256', secretKey).update(unsignedToken).digest('base64')
    const hashNoPadding = hash.replace(/={1,2}$/, '')
    return hashNoPadding
}


/* JWT 生成 */
console.log('*** Create JWT ***');

const header = {
  alg: 'HS256',   // アルゴリズム
  typ: 'JWT'      // Tokenタイプ
};
const headerBase64 = convertBase64(header)

const payload = {
  sub: '1234567890',  // subject (JWT の主語となる主体の識別子)
  iat: 1516239022     // issued at (JWT発行時刻)
};

const payloadBase64 = convertBase64(payload);

// 署名なしToken生成（ヘッダ と ペイロード を.で結合）
const unsignedToken = `${headerBase64}.${payloadBase64}`;

// 署名なしToken＋秘密鍵とHMAC-SHA256により署名生成
const secretKey = 'password';
const signature = encryptHmacSHA256(secretKey, unsignedToken);

// JWT 生成（署名なしToken と 署名 を.で結合）
const jwt = `${unsignedToken}.${signature}`

console.log('JWT: ' + jwt)
console.log('******************');


/* JWT検証 */
console.log('*** Verify JWT ***')
const secretKey2 = 'password';
// JWTを分解（署名なしTokenと署名）
const splits = jwt.split('.');
const unsignedTokenForVerify = [splits[0], splits[1]].join('.');
const signatureForVerify = splits[2];
// 署名なしToken＋秘密鍵とHMAC-SHA256により署名生成。
const createdSignature = encryptHmacSHA256(secretKey2, unsignedTokenForVerify)
console.log('signature in JWT: ' + signatureForVerify)
console.log('create signature: ' + createdSignature)
// JWT内の署名と生成した署名が一致すればOK
console.log(createdSignature === signatureForVerify)
console.log('******************');
