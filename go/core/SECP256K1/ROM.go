/*
 * Copyright (c) 2012-2020 MIRACL UK Ltd.
 *
 * This file is part of MIRACL Core
 * (see https://github.com/miracl/core).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* Fixed Data in ROM - Field and Curve parameters */

package SECP256K1

// Base Bits= 56

var Modulus = [...]Chunk{0xFFFFFEFFFFFC2F, 0xFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFF, 0xFFFFFFFF}
var ROI = [...]Chunk{0xFFFFFEFFFFFC2E, 0xFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFF, 0xFFFFFFFF}
var R2modp = [...]Chunk{0xA1000000000000, 0x7A2000E90, 0x1, 0x0, 0x0}
var SQRTm3= [...]Chunk {0x8D27AE1CD5F852,0x6D15DA14ECD47D,0xC2A797962CC61F,0x3507F1DF233770,0xA2D2BA9}
const MConst Chunk = 0x38091DD2253531

// Base Bits= 56


const CURVE_Cof_I int = 1
var CURVE_Cof = [...]Chunk{0x1, 0x0, 0x0, 0x0, 0x0}
const CURVE_B_I int = 7
var CURVE_B = [...]Chunk{0x7, 0x0, 0x0, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0xD25E8CD0364141, 0xDCE6AF48A03BBF, 0xFFFFFFFFFEBAAE, 0xFFFFFFFFFFFFFF, 0xFFFFFFFF}
var CURVE_Gx = [...]Chunk{0xF2815B16F81798, 0xFCDB2DCE28D959, 0x95CE870B07029B, 0xF9DCBBAC55A062, 0x79BE667E}
var CURVE_Gy = [...]Chunk{0x47D08FFB10D4B8, 0xB448A68554199C, 0xFC0E1108A8FD17, 0x26A3C4655DA4FB, 0x483ADA77}
var CURVE_HTPC= [...]Chunk {0xC813789E8624AA,0xCA45C23F508ECD,0x640A39CD8BBBFD,0x813FFE30F4D5B4,0xCCE8E9E8}
var CURVE_Ad= [...]Chunk {0x5447C01A444533,0xD363CB6F0E5D40,0x58F0F5D272E953,0xDD661ADCA08A55,0x3F8731AB}
var CURVE_Bd= [...]Chunk {0x6EB,0x0,0x0,0x0,0x0}
var PC=[13][5]Chunk {{0x38E38DAAAAA88C,0x8E38E38E38E38E,0xE38E38E38E38E3,0x38E38E38E38E38,0x8E38E38E},{0xCBD0B53D9DD262,0x6144037C40314E,0xDECA25CAECE450,0x23F234E6E2A413,0x534C328D},{0xFF1044F17C6581,0xD2FC0BF63B92DF,0xCEA7FD44C5D595,0xBC321D5B9F315,0x7D3D4C8},{0x38E38DAAAAA8C7,0x8E38E38E38E38E,0xE38E38E38E38E3,0x38E38E38E38E38,0x8E38E38E},{0x2A56612A8C6D14,0x6B641F5E41BBC5,0xD51B54225406D3,0x4383DC1DF7C4B2,0xEDADC6F6},{0xE6B745781EB49B,0x409542F8487D9F,0xCBB7B640DD86CD,0x3D94918A9CA34C,0xD3577119},{0xBDA12F38E38D84,0x2F684BDA12F684,0x4BDA12F684BDA1,0x12F684BDA12F68,0x2F684BDA},{0x65E85A9ECEE931,0x30A201BE2018A7,0xEF6512E5767228,0x91F91A73715209,0x29A61946},{0xFC90FC201D71A3,0xB046D686DA6FDF,0x4B12A0A6D5647A,0xD5CB7C0FA9D0A5,0xC75E0C32},{0x2F684B8E38E23C,0x4BDA12F684BDA1,0x12F684BDA12F68,0x84BDA12F684BDA,0x4BDA12F6},{0xBF8192BFD2A76F,0x21162F0D6299A7,0x3FA8FE337E0A3D,0x6545CA2CF3A70C,0x6484AA71},{0xB425D2685C2573,0xC1BFC8E8D978DF,0x632722C2989467,0xB8BDB49FD5E9E6,0x7A06534B},{0xFFFFFEFFFFF93B,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFF}}

