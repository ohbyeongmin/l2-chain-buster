package optimism

var (
	privKeys = []string{
		"0xeb0e6097d5a5d630bcc66780d1fce68b759d8e3db2ab1566ed43657970d564a2",
		"0x430e2662446e419d68b2b4b7545e0be3ab5783b4bed4e91a7327f0aba679d773",
		"0x56f2ef5697e3630213d68fd1a40db707d986297b061f8767256bbd1f48516361",
		"0xc149e722902f14cfe1b549572d64fc74e614d24c3aea7af35f051a6469f3b2b1",
		"0x4237f9a2d55f846a3bbc716245974b1fb46c0712f17db9d933cde72f03f4723e",
		"0x0a5dc8a63e4739695294c03d069701caef51d1356d938c3d37ab53e7d337a5c8",
		"0x0dc2b28f8d4cf51db306290dfdacdd7acec3733c47d8f5a99e73f7c4c3111e3e",
		"0x32075254dd8d412e23a05856bb2ac7978e3d9db79117b92b26b7802008f39f36",
		"0x5298fdf5e89ea67149762964162af72cd2708cc2cf84e5dc202507a19dabd1a5",
		"0x01c95e7838114cd78dc12467f6e5f8306a0e0e283f9f767ed02844e27408f17d",
		"0xd49d6dc532b470896f32b0bc308d80ee722fe74a6415ccdf3415aadd56e6530d",
		"0x1317b72f84661e524d61f12357c07cd16ddff5671e5332764d132113bb13c95d",
		"0x1b140d5cc324c8b9352b60771bcbe16b14cd061e658f9ed86984c33d395f33db",
		"0xd25160b18cb264b86b62da9ee6a30643ffba698891a60364d2b76dfacdbd1c47",
		"0xca0cd4c0dcecac62b24061b3f0a62e1f2d74cde89d0dca7ac10af1370b66916e",
		"0x0bf35f91204ab0942d49e8e7c96c84e197904dfcb7cf4ec4f95ee941123f30e0",
		"0x4c1a81e5693594c0ff5b99a4a33eca23bc7e116766be0ac61554656da82cb89a",
		"0x89f8b823ede3ba77af77f791726e10f93035bcd613dad468899d088e165970a7",
		"0x059132f3c232644ca2d34f6bad1fbcd7b57625dac987f244b57e1ecc74912359",
		"0x2a6377febfa11cc90d30bb486cb041aadc0090cb6541c6aa487a53d83bce95f7",
		"0x3da1c084e1afa57eb17f084afac91bca31e7726abeb9d38e3e8eb3b18a6b2e57",
		"0x416602b7f1507120959ffc918c827db37ce69839bbfc9f27ef069bf5b44dacc2",
		"0x5c6065e69c5a9a2c996863ed8127fdd31e95c65bfcf62a87e243c0559f0639cd",
		"0x96196a13ef2776b1e829266becefdf9a7d374302d7e7cc023a83184d44b21c3c",
		"0x17821d2510fada54178248a9bdeb41175c7c7cdefb9b52777de566c62479bbe4",
		"0x2a6e5c134147ae3f1e5e1546129925f8006f5153f925c3f7b786b9eaf5179a53",
		"0x5fe70255e1d7fb680d4e8912dc0daf7706488967963f4cf4e8d1d92cc5af2b8b",
		"0x54fcf6a2a1ba3d4de8f9333372ebebd65151561cf53a301ece577a4ccde6e23b",
		"0x21cc78018560203501465c41388b4fe099f4bf6581878882fbc77beac2956cb2",
		"0x56f74f58d11dad5bf159ad0fd4919dfc2a0ef240fb53c2361a14e671b08e71ad",
		"0xfd09c5707784085363e9135852ed7ce94adc179d4a3f7420a22158aa99ec4486",
		"0xf1b2aa06d5f148d28f83745f3372a8d5cc946f3362f951342966f6f1dc2ed42c",
		"0xa5122d377ac6f78318c988a180b94bc2c40b113dd4fe48a9395c6fdd48eae385",
		"0x213d191f1d696d879e65e3dd2dd7259d66cf3802de161271c8602b716e0bf97f",
		"0xe9a10a2551d411f5e6fa8f04e00efd729403e181b824ebd0d03eac2f9448b3c2",
		"0x5edc7ffc17b14f7464ef8654ac7f415ccec80ce95de43d9a155a4187e53e0599",
		"0x1f8b57bef2a85db6d82cea3dc0c3eb5d0bc2127c6bc5c1210fa3f8fb86f608c2",
		"0x7057d09d85d22ba49b48324a0477052fb2971f2171cc6aa230559fc52315648d",
		"0xdc3f6e7bfef79a1025dc0f136b741f43fdbaadc76e5d57e6bd855b8bce8edd32",
		"0xe18c1e1044393a24dd03f1f02f2dd8f601a97da7e6d9e3b440fc260724255e2b",
		"0x5acb882288e9f5377fcc008af408e02cfd9bc14b7708a4c7606cc7e7fda63fa9",
		"0x8c2265fe44a29d384cf575c05a9ade451bde6e81d22d4458d6b2f7deb18fcddb",
		"0xff81a7d9fdd54af94e9c0000434d1429101faec1b6be9a3d29fff42fe708aece",
		"0xab3388ae09550538a0fea49f87a19f43cf39010c4df2d15fe57bf3e3ec83fcb5",
		"0x11f8356b6c0030aff5121c9aaf6e4a765125f8168cbb674fed32627d3f6120de",
		"0xf18a646712a0c1243769386e3a8f765b5f08a43f2513052e69ce14445965879d",
		"0x394ecf44233ce097b81d51f0e748ccc0989ba954d766b697665f4368363ee8c5",
		"0xee8c11902e21f2756aa35ee3979017cf308a785d1fc4a1cf623a21f10077b7c1",
		"0x79ac750a5881a46259e949ba53a972834cdb6c3d0c8ab557c998f6a27a87323b",
		"0x776e88b8903b90bea2d1b3695ce1a28849d3f94660a3b23444816063cf471f57",
		"0xb961c334cb6b96e559218a10dc330531acb88e935fea454ca2117c051da27f4c",
		"0xd47c8ba92ff61fc16e1884a30b338d05ab794299efa46c9faad3d91ab1804312",
		"0x48eca7fc7b028ba17d0e065abc4f0a6a7fcd407cfc28ed1d10bc177c8a35a7c9",
		"0xf6ffcb8927cfbb1a7d4538992477baf6b0cda54559959c3540a04e645fb2ea47",
		"0x99bcfd4f1ce063ed7f7f08006bccab0951d1fd4d2771c519948c6643d4382e38",
		"0x977a1703d542e29e55858d280598edca6938c0bf6ea53aefe4a0994f3cbe35fc",
		"0xa01dc4eaaaf124c578ecf142828d89112ac50d2f696a62eb3f8e89cd0426c4a4",
		"0xbdca74323d3ce8b5069ff5e19e183b5967fc06fed8c745b67c1fccbb1fe50917",
		"0x2994f655f0965b03ec7d7be7498840b879c6f110378bf99b71f643eb29917f7c",
		"0x561777e953193ed402bbcc2eed8b88ecad34c5fef2d6185e63f9afbe54b9c3b6",
		"0x7e2f55cce87f9dae4f435abb913a80cd73fe2171967c9a530fe30d0670e0855f",
		"0xd2fabde7e3f20ba582ac9fe0a5f0e3e8aef9cd8e5e5a2f5df1764b7cb583d3f9",
		"0x228d1cafc37833e263ebc753d14fab3620f7b3b1b144db5ef86d1c5f4ddfe47b",
		"0xa4d27c8a5a62d9d681321dd2fb781593833c7e5645ecbc1c6b14329f381e4362",
		"0x38f59871ad39e23e207db55297e950c491792eec9ca2b0b8213a44f55bf3a972",
		"0xfd2f61194ed705ca8b6b50cd797ca3ab1423b335949e451d07e04db60993e80a",
		"0x27297a2edf7d09f3204f7a0d753002dbaa2fac85405dd0fa93afa63f4fd9fe49",
		"0x692045d29e66117bea3d684ac96dcfa49683cadf503b15ed9c5fd2770cf75a49",
		"0xb9534b638c9ce79a104028df2e77185f94b6a8f0bc2a508fe6bada1d0e9486b5",
		"0xba4864cf70843455fab73655233f7d675617f5c2e52a36794d480096858452c9",
		"0xe15580731afe09492d8834a9212286a3cc8b1077241ebcbfec0226c41d8e8122",
		"0x670719cfba6457031a43c72a8c99fe07a27f5d3eeb47e848f38719c5625b14f6",
		"0x909a1fdf7938e4a51a39e79751797198885f582b0748d3f9ebac5de0ac0aaf75",
		"0xd40ad3458bc3f54ece059f6e6d924fb65e7110201216efd3b2d34069911da7e3",
		"0x7714c79ddedd971d00ce57b095e3008bc853f149baa4a4dbd70bc854c9f6a88a",
		"0xcfde9448687226bfefee74ff0772875116f798d4e8120540195ad929f6e706d6",
		"0xa6a3930da6ea38871cf53854cbc1b09d5e0fca6313cda54335f71ec1f2d78f0b",
		"0x4f5cf61d34c2c90b06225955dbbef2a46066f4033523414e40f90611a2ded4df",
		"0x20c98e0d977ac9c58feb4654f78252fdc3f10cc08bbf6d46c486a88bfd7a48f3",
		"0x19ac3f1853966a2dc0ccf0d65534a2bb826e2c7ce41facc4549da5c5cf39e19b",
		"0x43e4bc1d00abfe7b0f4945dbf0c1b650d924e24cb57c249f26bfea23c8945b62",
		"0xa204c59a6473fe5f401aaed00a6890d74e4a816b48e90bd70dd8f2c72dc23cdb",
		"0x9a97a3f924ce378e37cc8314905d71ed959a63d203d18faf2c9a4e57e4990671",
		"0xe762ea282fb269920aa304311c669537c8e8909f94472345545e1e211c6b0d77",
		"0x36e7a3b9bb194070478a5e149d6dd462c7178dec277b6385b5f4d546e5301077",
		"0xa9d6fb604034ddb28af2500081c1cc2e62f2257a385de00b94e6152e0e40cecc",
		"0x1f1aad895823fe5de4b5c7b2db05b555f30d7794b0947ed9a734f3fe8d0da919",
		"0xbf7142732b87230d33326387d5e0cdc3254ceb21b77b92486d20565051fa1486",
		"0x99d7b30644856b7685ec62297c90e604beafc1598d5080e5982d14c4743d85ee",
		"0xd5a8fdbbfebceae7ec2dfed6fcf44730e518d480c752d9fdd2be07a18cfd402c",
		"0xb5f72f4ef04eaa247049170df58340b3a426bbf224c138eb89efac1bf970fdd9",
		"0x9c231f0cc5a2aa6710aba80e6bf4124ffe10c39fd75806c5def344cb0f874827",
		"0x9ea7c6697a06620973b5418a36f379a8eab896548165f7ee836f0939103ee3a7",
		"0x9e8861f21c9a814aece0e810830f647c5198a389aff16bca9db1cb3dcaf58d44",
		"0xbf9f742d38f76f612fd44451a02c69cfa6fa8f7942869e085200d25f204c8caf",
		"0xd473184a99dd4e8ec7448d49c7a955892bb438535a1569ca6058b7245d1cdfad",
		"0x64a3cb328be20ada09b7822cd2ab6016bb38ee37d1aba3e8d279141f5995c6af",
		"0x2980309c5da0da4146d7bba994c111165d904cd5248f5b502793d21ec4cd4fd4",
		"0xc95856578699ff3edab6300894417e9ed8744d7f1df8a6596a223ca91b9d2ca7",
		"0xb0788eef8964490ce6b102adeb226b302d32987b72f6d65af72ddd62a152a6ff",
	}
	addrs = []string{
		"0x8A34C56faB841929326563A087455136Ac8E4126",
		"0xd8a30f508b0CE827e25815F199eEf4508bAc97e1",
		"0xfbb6c04325eF1450b541f9451c813E49B77C63c4",
		"0x6882aCA825a5d506bDc4Ace832b094a844C1839B",
		"0xB19C9c25C1Cc4841bEd9Cf9ac05c9C5cb4925e9D",
		"0xba209B2c377968B822a1075A3f6703406f639fdf",
		"0xA232cEC38b6cD7b7712137b7151AC0E96e425405",
		"0x6337c1ca6f87F0a802D792730fbcCdb5b8431034",
		"0x2Ec51668EB26eF5537E1259c0456aF65D261Cb6d",
		"0x3780B4559E5F03F9884E1BEDdE6b5F5258e51Cbf",
		"0x2c25e0b8dDAE91f60107b37aa9f62db45AEE88B2",
		"0x4e42E6eFf4584CA92e1dc7121DaBCA4eE458B1b1",
		"0xca8EB41049F6E130D006F9E94a5b6C55D03b4bF9",
		"0xcE07B903A57D2AbfE4755a18F77e975d1031F1Ed",
		"0xC746125c7F52FCa3De7C1D72Cd3fE8C6Ba6757Ce",
		"0x65b24ff021275E8e9FAB42866d2f0030b9f816Df",
		"0x443D7dAdb1aDf6b63290A4D38D960b5cB131ce84",
		"0x4257b137AE582F75D6b96149aA1B684AAd4667B5",
		"0xa6ccDffC3947Fe2A154f0a3A51b99Afc5075046a",
		"0xE9e305684e1806736efBcb6e9859b01a302e38FB",
		"0xc3CD48AA99fA35F2cf85cd7a341d2548532eCb37",
		"0xC3846ef419cfE85D8B967a79b55dd215eea240A2",
		"0x42330CB450c18a8E9eb3dcAb8541B8d47c16a7C8",
		"0xCFAE3Cf25beca4ca323E0E0bb995374EbaC02089",
		"0x3BEe9f536Ff3732fa91761fe070f990Bd1F86fE6",
		"0xe97309B35Ef0dB2E9E941eD9E0898EAAE18713f0",
		"0x11389b78881DCA680A72e4BF4EAFB3c437bB2eE1",
		"0x4d71da4D62b778D792A53a203F01a5037b286df3",
		"0x57934459c33B4B5DE293A184cBE8f8303B633f93",
		"0x7057Eea59A1E87661e431cE93a875310f5f5a2D1",
		"0x38102632365eb0D9484DbE7b6F8eeB86C18E84a4",
		"0x90C1A91d8C8E9f9CfEae1C7230A8ab27c3A90532",
		"0xf7bB6E9A22e383Df471Aa40df332d955c721Ed9E",
		"0x4bB9C237634b5d7d3d15A65AcD91073baf7398f9",
		"0x8F41cD37F283D6E166d1340C97282c901eD03b5e",
		"0xfc2B942a9604B133C6747E881e7E4BBE6Aba7C53",
		"0x04030c05C2d81f40D1EdE7f2D3451aa7D75bB31E",
		"0x6dD2D6effB6166E5B1f3204254ec11c358F782D7",
		"0x94bf959e0871bFc93E0151381aabFE4AB6D6D221",
		"0x97937386faeE97d209Fe5E3a3DF494782D7FaF10",
		"0x7F9Da93c2a207b752799B83c5ADB029f0afA898C",
		"0x5f609848Ca12749F2E5F3D781397c6e1c177BFaD",
		"0xB2Fb31e5A9bD959bCF2B219bee1Def08B6D31C91",
		"0x82fCcf31A2140151170F94961575368F0351DDf1",
		"0x1eE8C6B45de10b97656d9625DA9E17e1b82f5d72",
		"0xf879234CBd8c85F8B2BF7c06C61E158693F8D82B",
		"0x5e4EEFCD2Cf73BD22f2113d68cFBd130389b3502",
		"0x0F301F1af9d6d380eb3B8C9402276B1c769e9f60",
		"0x20ce89f88556684D7ea6fbCC5621e1904ea2A0B1",
		"0xb222b5BBdBa008bf09E3063dF52aFe36E585d25e",
		"0x1a1468fC7E19BFf7a1aE9171df380bD5dd4d7574",
		"0x59DC00d20eAB3A8170cAd53919ABAD57E921fdD0",
		"0xF15C77A42414AEFFB4D531Ba722D5662F5606C27",
		"0xeAd83D1a621774238653d7eF2CdDE0bd1D4Be73F",
		"0xD64796290110Aa71A6613bC7E7879b44eAeAf5C1",
		"0x34CDAd9922fdA5720A4dD6B714e0604f2871d2C2",
		"0x6ce3113Ed013e66f60CFA38465fE637FF86eFCE6",
		"0x518Bf430eEf0B33bA1fcd4D54B9A5e330Cf82656",
		"0x7cf32756ea37560147538773E3e1467122e2F370",
		"0x4E219A968335BF4F266f454251D19161f31d466E",
		"0x7CCA32ad755f55A83be6Ffa46721F7a9B3C3f5c7",
		"0x2267f47B6efD7D61eA76CBd1184e65872550a5E2",
		"0xC1555D3A00D5B92e242553c3dd8ac7c26d0946cA",
		"0x522Afde67F581e83C20597152B5f14A1b6065249",
		"0x508C7AE1a78eDe80815D7c87c784D7F3C0F47d95",
		"0x09077CB17fbEd89075642aB584DdD3a5407344d0",
		"0xbe3C2cc99DF3212d0877FA48c06519FF3b673cFE",
		"0xCf878F3B8bC5C188B05F51a985dD8Ebd64f8CDb3",
		"0x795870bf3E53E6C436Fe28650E5A4c5CEC228D10",
		"0x2739E685424bE3f3A3a43FA918bb1E56a1Ba6d28",
		"0x0Bc82AcEC8c075d2Ad294C79c0e6A803a564F121",
		"0x8DB771cEfa926c88328248B3EEF5DE821d084350",
		"0x07111DF3E9cA9980245c9c5df9171FC462e9F8fb",
		"0xc1bAA1835D1e62322e362E6D59E77b8b30CC3b82",
		"0x9BA8C4F23748e82d97052514cbe80C72Dac0FE1D",
		"0x45AFA7dfFbDb56985c1BFD4EE3435E515B164CD3",
		"0x658C9bF16A14eb1561DccA8C0cbfFC8c473CBA03",
		"0x938968bF21faD8050F364e2E299129a4C97128F8",
		"0x07EBFbFE4B8797439126570F804435CEba4F75A7",
		"0xc5022A1330142AdfEF5F8A7D4D3581cd27A8BFdE",
		"0x9049ae343294227B2D875cf582fF835802c048Cd",
		"0xFBE186C4B1AF30c80c5BEc396fbc0bd8CC8fA4a2",
		"0x073f745ab9dBc72B879956D96D9B3567b6A7b70E",
		"0xCED11637DEb2e60c3c8AE7bFCB20FAa7b849d9a8",
		"0x5fE6fbf89A7cf3Aee5fA92504cd5B1805F6b16b6",
		"0x99e6BDe7a47826C6392e2E4b8d31e6da96b48494",
		"0xB628e4a151cFAB6602A61c34c65b878eC76df3B0",
		"0x9f39579aABEe32092cB0Dd70C63652025730e5f1",
		"0x96921c7a54c630C9915E6021E6d0181FE123FC15",
		"0xB919000b961ca67e5688C88032Ff023a5B6d7d71",
		"0x51aDe8aad420d02A19eAb13136bCe28A3440708a",
		"0xE77b8226d393917DcbD1383f12cB86323F76eAa7",
		"0xd19F054Cacca52f4706669025B785D63382b78D6",
		"0xAEaC649FA24A4761A5d67c499C522dFE09706931",
		"0x51027D08851F4E5aDDcf5181bf70529554F0c36e",
		"0xDF7beDca8F60499f2D7D5868a2b9a443f6ac1eb2",
		"0x593461023946C68DFfB82928A53E594580418Da9",
		"0x76ddfdFe37e244F9355648c287Be07C66F793484",
		"0xB1e3e27AbfF20dE2038734C1604DB9cF9913aC59",
		"0x7ABa0EbbDDb708022A3D5E7471D05f48DA563C3d",
	}
)