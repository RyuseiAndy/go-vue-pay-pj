# マイクロ決済サービスを利用したショッピングサイト  :dragon: #
<p>
<span class="mgr-10"></span>
</p>
<h3 align="left">概要</h3>
ー  gRPCを利用したマイクロ決済ショッピングサイトです。簡単なショッピングサイトで購入したいアイテムを選択するとアイテムの画像、値段、商品名が表示され購入に進むと（VISA、JCB、クレジットカード）の番号を入力する事で商品の購入をする事ができる仕様になっています。
&nbsp;
<div align="center">
<img width="1434" alt="スクリーンショット 2022-02-01 15 26 28" src="https://user-images.githubusercontent.com/83407832/151931139-64938e51-b4fc-490b-8794-2abb75590e5f.png">
</div>
&nbsp;
<h4 align="center">現在のホーム画面</h4>
&nbsp;
<div align="center">
<img width="1425" alt="スクリーンショット 2022-02-01 15 26 42" src="https://user-images.githubusercontent.com/83407832/151931145-df0f06fd-9ca8-4784-b956-d250beb048e3.png">
</div>
&nbsp;
<h4 align="center">アイテム選択がめん</h4>
&nbsp;
<div align="center">
<img width="1158" alt="スクリーンショット 2022-02-01 15 27 07" src="https://user-images.githubusercontent.com/83407832/151931150-39a20d67-b9d8-433b-a0dd-aaf4d7f94726.png">
</div>
&nbsp;
<h4 align="center">決済画面</h4>
&nbsp;

<h3 align="left">使用技術</h3>
ー  gRPC、 PAY.JP, バック側：Go(フレームワーク Gin), フロント側：Vue.js
&nbsp;
<h3 align="left">背景</h3>
ー  作成しようと考えたきっかけはECサイトを利用した際に,ECサイトがどの様に構築されているか気になり自分の手で作成したいと考えたからです。
ー  自分の友人が私の好きなブランドサイトのショッピングサイトの構築に携わっているのを知り、私もWebエンジニアになりたいと思ったのがきっかけです。
&nbsp;
<h3 align="left">実装機能</h3>
ー  アイテム（商品）の表示：名前、値段、アイテム画像
ー  商品購入時の決済機能
ー  これから実装予定の機能 :dragon:
ー  ユーザー個人のページ  
ー  ログイン機能、ログアウト機能
ー  購入履歴閲覧ページ
&nbsp;
<h3 align="left">工夫</h3>
ー  最近ではWebベースのシステムにおいてHTTP/HTTPSベースでサーバー・クライアント間のやり取りが行われ、その際のデータフォーマートにはXMLを利用する「XML-RPC」や、同じくHTTP/HTTPSベースでデータフォーマットにJSONを利用する「JSON-RPC」が多く使われている。しかし、基本的にテキストベースで情報をやり取りするためデータの転送効率が悪く、バイナリデータを扱いにくいという問題があったためそれを解決するために今回gRPCを利用した。
ー  gRPCはサーバーおよびクライアント側のコードを自動的に生成するツールが提供されているため、これを利用することで簡単にサーバーおよびクライアントを実装できるようになっている
&nbsp;

<span class="mgr-10"></span>

 &#x1f4a3;  Still working to complete the project.  &#x1f4a3; 
