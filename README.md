# マイクロ決済サービスを利用したショッピングサイト  :dragon: #
<p>
<span class="mgr-10"></span>
</p>
<h3 align="left">概要</h3>
ー  gRPCを利用したマイクロ決済ショッピングサイトです。簡単なショッピングサイトで購入したいアイテムを選択するとアイテムの画像、値段、商品名が表示され購入に進むと（VISA、JCB、クレジットカード）の番号を入力する事で商品の購入をする事ができる仕様になっています。
&nbsp;
<h4 align="center">現在のホーム画面</h4>
<div align="center">
<img width="980" alt="スクリーンショット 2022-02-01 15 26 28" src="https://user-images.githubusercontent.com/83407832/151931139-64938e51-b4fc-490b-8794-2abb75590e5f.png">
</div>
<h4 align="center">アイテム選択画面</h4>
<div align="center">
<img width="980" alt="スクリーンショット 2022-02-01 17 00 56" src="https://user-images.githubusercontent.com/83407832/151932292-dc87015f-7465-40d1-a9cc-d0ca3d0343b4.png">
</div>
<h4 align="center">決済画面</h4>
<div align="center">
<img width="980" alt="スクリーンショット 2022-02-01 15 27 07" src="https://user-images.githubusercontent.com/83407832/151931150-39a20d67-b9d8-433b-a0dd-aaf4d7f94726.png">
</div>
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
<p>
&nbsp;
ー  これから実装予定の機能 :dragon:
</p>
<p>
ー  ユーザー個人のページ  
ー  ログイン機能、ログアウト機能
ー  購入履歴閲覧ページ
</p>
&nbsp;
 
<h3 align="left">工夫</h3>
<p>
 ー  最近ではWebベースのシステムにおいてHTTP/HTTPSベースでサーバー・クライアント間のやり取りが行われ、その際のデータフォーマートにはXMLを利用する「XML-RPC」や、同じくHTTP/HTTPSベースでデータフォーマットにJSONを利用する「JSON-RPC」が多く使われている。しかし、基本的にテキストベースで情報をやり取りするためデータの転送効率が悪く、バイナリデータを扱いにくいという問題があったためそれを解決するために今回gRPCを利用した。
</p>
<p>
 ー  gRPCはサーバーおよびクライアント側のコードを自動的に生成するツールが提供されているため、これを利用することで簡単にサーバーおよびクライアントを実装できるようになっている
</p>
&nbsp;

<span class="mgr-10"></span>

 &#x1f4a3;  Still working to complete the project.  &#x1f4a3; 
