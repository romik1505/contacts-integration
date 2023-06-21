package model

// Параметры запроса колбека от AmoCRM в наш сервис
type AmoAddWidgetHook struct {
	Code     string `schema:"code"`      // Код авторизации
	State    string `schema:"state"`     // Состояние при открытии окна
	Referer  string `schema:"referer"`   // Адрес аккаунта пользователя
	Platform int    `schema:"platform"`  // Номер платформы на которой установлен виджет 1-amocrm.ru, 2-amocrm.com/kommo.com
	ClientID string `schema:"client_id"` // ID интеграции
}

// GET /sign_in
//?code=def502006df60b71998fcb56b8997cf3048d5de9d684dd4e936d2943b5e8ce5b65f48e22af7a792fc6669d8ed8988074e6218fc9e00dfc1673f197de2de68800888887032361111eae7cde2e386538511e05fc11e61c56d72cd6b4b9be2c59a526f171e507e291ae8f34f16027c2a8b2dad8df0e85e5bb4f11c366137de56d5aafd257da68c294089d57f3cd2f954183d61547a4175c110299f939754aefdf2aa44e1b8f81674c2a1915b6d2f83337ecc4c1ef244d9c12e8df01d23900690bec914ef6fd0c3463d0f35d4d64b3df8233ba376f1e16695f29f481d1d85b5d8e7e699385d69034116ec22cee7ee2d7de369bd654de314679e4803406a716c2fbf64ae1f462384d335c154ce6cbe9f8d09ecc1092794cbe509a587e4ff88fdd67b9acbe0b7396022714c3bcf51229c2dfd45eb8c0d1d8ad75e95ce7644691eab9c8338f9e473b0f7e61c19f0d5f71dd111d028dd1c51c3636a20ed23deade2e4b9451a64aa39d6cc95bc588d4e3f03f3c181cd9711dca54b0cfedb8a24a95334ec01303a4508db6d2d90fa7e7ca95d219ff79a657844b444104607fe13da29072077a4dab17a4bec9a156665098be21b633ba6ee4dea6e2591ed5ca3ceed677df20b84a78226ce61333423cb76da0a0cbbf9c44ec32520aff3d0c47e4defd735e7d58a9692832dfd7070c8c8b35ee5ddf29c661f2c9d0dba6e2539748b860b80f18b757ee
//&state=1
//&referer=droman.amocrm.ru
//&platform=1
//&client_id=3db50128-f43a-4d32-bcce-2fef22baf758
