function addTask(text) {
    $("#task_list").append(
        $("<li>").append($("<span>").text(text)).append(
            $("<button>").text("Удалить").click(function() {
                $(this).parent().remove();
            })
        ));
}

$(document).ready(function() {
    $("#root").append($("<ul>").attr("id", "task_list"))
              .append($("<input>").attr("id", "add_task_input"))
              .append($("<button>").attr("id", "add_task").text("Добавить"));
    $("#add_task").click(function() {
        addTask($("#add_task_input").val());
    });
    addTask("Сделать задание #3 по web-программированию");
});
