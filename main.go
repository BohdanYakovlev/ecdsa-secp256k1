package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

type point struct {
	x *big.Int
	y *big.Int
}

func addPoints(p1, p2 *point) *point {
	res := point{big.NewInt(0), big.NewInt(0)}

	deltaPx := new(big.Int).Sub(p2.x, p1.x)
	inverse := new(big.Int).ModInverse(deltaPx, modN)

	sCof := new(big.Int).Sub(p2.y, p1.y)
	sCof.Mul(sCof, inverse)
	sCof.Mod(sCof, modN)

	res.x.Mul(sCof, sCof)
	res.x.Sub(res.x, p1.x)
	res.x.Sub(res.x, p2.x)
	res.x.Mod(res.x, modN)

	deltaX := new(big.Int).Sub(p1.x, res.x)

	res.y.Mul(sCof, deltaX)
	//res.y.Mod(res.y, modN)
	res.y.Sub(res.y, p1.y)
	res.y.Mod(res.y, modN)

	return &res
}

func doublePoint(p *point) *point {
	res := point{big.NewInt(0), big.NewInt(0)}

	xPov := new(big.Int).Mul(p.x, p.x)

	trplXPov := new(big.Int).Mul(xPov, big.NewInt(3))
	//trplXPov.Mod(trplXPov, modN)

	doubleY := new(big.Int).Mul(p.y, big.NewInt(2))
	//doubleY.Mod(doubleY, modN)

	inverse := new(big.Int).ModInverse(doubleY, modN)
	sCof := new(big.Int).Mul(trplXPov, inverse)
	sCof.Mod(sCof, modN)

	doubleX := new(big.Int).Mul(p.x, big.NewInt(2))
	//doubleX.Mod(doubleX, modN)

	res.x = res.x.Mul(sCof, sCof)
	//res.x.Mod(res.x, modN)
	res.x.Sub(res.x, doubleX)
	res.x.Mod(res.x, modN)

	deltaX := new(big.Int).Sub(p.x, res.x)

	res.y.Mul(sCof, deltaX)
	//res.y.Mod(res.y, modN)
	res.y.Sub(res.y, p.y)
	res.y.Mod(res.y, modN)

	return &res
}

func mulPointOnScalar(p *point, n *big.Int) *point {
	if n.Cmp(big.NewInt(1)) == 0 {
		return p
	}

	if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return doublePoint(mulPointOnScalar(p, new(big.Int).Quo(n, big.NewInt(2))))
	}
	return addPoints(p, mulPointOnScalar(p, new(big.Int).Sub(n, big.NewInt(1))))
}

var basePoint *point
var exponentN *big.Int
var modN *big.Int

type signature struct {
	r *big.Int
	s *big.Int
}

func setBaseValues() {
	basePoint = &point{big.NewInt(0), big.NewInt(0)}
	basePoint.x.SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	basePoint.y.SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	exponentN = big.NewInt(0)
	exponentN.SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	modN = big.NewInt(0)
	modN.SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
}

func getPrivateKey() *big.Int {
	temp := new(big.Float).Mul(new(big.Float).SetInt(exponentN), new(big.Float).SetFloat64(rand.Float64()))
	res := new(big.Int).SetInt64(0)
	temp.Int(res)
	return res
}

func getPublicKey(privateKey *big.Int) *point {
	return mulPointOnScalar(basePoint, privateKey)
}

func getSigECDSA(message *big.Int, privateKey *big.Int) *signature {
	res := new(signature)
	res.r = new(big.Int)
	res.s = new(big.Int)
	k := new(big.Int)
	for {
		temp := new(big.Float).Mul(new(big.Float).SetInt(exponentN), new(big.Float).SetFloat64(rand.Float64()))
		k = new(big.Int).SetInt64(0)
		temp.Int(k)
		mulGk := mulPointOnScalar(basePoint, k)
		res.r.Mod(mulGk.x, exponentN)
		if res.r.Cmp(big.NewInt(0)) != 0 {
			break
		}
	}
	inverseK := new(big.Int).ModInverse(k, exponentN)
	temp := new(big.Int).Mul(privateKey, res.r)
	res.s.Add(message, temp)
	res.s.Mul(res.s, inverseK)
	res.s.Mod(res.s, exponentN)
	return res
}

func getVerSigECDSA(publicKey *point, message *big.Int, messageSig *signature) bool {
	inverseS := new(big.Int).ModInverse(messageSig.s, exponentN)
	u := new(big.Int).Mul(message, inverseS)
	u.Mod(u, exponentN)
	v := new(big.Int).Mul(messageSig.r, inverseS)
	v.Mod(v, exponentN)
	mulGu := mulPointOnScalar(basePoint, u)
	addPv := addPoints(mulGu, mulPointOnScalar(publicKey, v))
	temp := new(big.Int)
	temp.Mod(addPv.x, exponentN)
	return temp.Cmp(messageSig.r) == 0
}

func main() {
	setBaseValues()

	privateKey := getPrivateKey()
	publicKey := getPublicKey(privateKey)

	temp := getSigECDSA(big.NewInt(2012), privateKey)
	fmt.Println(getVerSigECDSA(publicKey, big.NewInt(2012), temp))
}
