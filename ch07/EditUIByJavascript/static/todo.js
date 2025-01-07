// defer 属性によって、DOM の構築がすべて完了してから setup 関数を呼び出す

setup();

/**
* 各 ToDo を表示する input 要素すべてに、click イベントハンドラを設定する関数
*/
function setup() {
    const elements = document.querySelectorAll('input[type=["text"].todo');
    for (let i = 0; i < elements.length; i++) {
        elements[i].addEventListener("click", onClickTodoInput);
    }
}

/**
 * ToDo 項目がクリックされた時のイベントハンドラ。
 * @param {Event} イベントオブジェクト
 */

function onClickTodoInput(event) {
    const todoInput = event.target;

    // input 要素を編集可能にする
    todoInput.readOnly = false;

    // save/cancel ボタンを囲う div 要素の取得
    const todoEditorControl = todoInput.nextElementSibling;
    
    if (todoEditorControl.classList.contains("todo-item-control"))
    {
        // div 要素から hidden クラスを取り除くことで、ボタンが画面に表示されるように変更する
        todoEditorControl.classList.remove("hidden");
    }
}
