(function() {
    usp.namespace("usp.admin.menus.list");

    usp.admin.menus.list.add = function name(params) {
        window.location.href = "Add";
    }
    usp.admin.menus.list.edit = function name(params) {
        if (usp.admin.menus.list.table.row('.selected').data())
            window.location.href = "Edit?id=" + usp.admin.menus.list.table.row('.selected').data().ID;
    }

    usp.addmin.menus.list.remove = function name(params) {
        if (usp.admin.menus.list.table.row('.selected').data()) {
            $.ajax({
                url: "Del",
                data: {
                    id: usp.admin.menus.list.table.row('.selected').data().ID
                },
                async: false,
                success: function(response) {
                    if (response.status) {
                        usp.admin.menus.edit.viewModel = ko.observable(response.data)
                        usp.Notify('系统信息', response.info, 'success')
                    } else {
                        usp.Notify('系统信息', response.info, 'alert')
                    }
                }
            });
        }
    }
    usp.admin.menus.list.table = null;
    usp.admin.menus.list.init = function(id, url) {
        ko.applyBindings();
        usp.admin.menus.list.table = $(id).DataTable({
            iDisplayLength: 10,
            "bServerSide": true,
            "bPaginate": true, //是否允许分页
            "bLengthChange": true, //是否显示每页显示条数
            "bFilter": true, //是否启用条件查询
            "bSort": true, //是否启用列排序
            "bInfo": true, //是否显示分页信息 
            "language": {
                "sProcessing": "处理中...",
                "sLengthMenu": "显示 _MENU_ 项结果",
                "sZeroRecords": "没有匹配结果",
                "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
                "sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
                "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
                "sInfoPostFix": "",
                "sSearch": "搜索:",
                "sUrl": "",
                "sEmptyTable": "表中数据为空",
                "sLoadingRecords": "载入中...",
                "sInfoThousands": ",",
                "oPaginate": {
                    "sFirst": "首页",
                    "sPrevious": "上页",
                    "sNext": "下页",
                    "sLast": "末页"
                },
                "oAria": {
                    "sSortAscending": ": 以升序排列此列",
                    "sSortDescending": ": 以降序排列此列"
                }
            },
            ajax: url,
            columnDefs: [{
                title: "ID",
                name: "ID",
                targets: 0,
                orderable: false,
            }, {
                title: '图标',
                name: "Icon",
                "targets": 1
            }, {
                title: '区域',
                "name": "Area",
                "targets": 2
            }, {
                title: '控制器',
                "name": "Controller",
                "targets": 3
            }, {
                title: 'URL地址',
                "name": "URL",
                "targets": 4
            }, ],
            columns: [{
                data: "ID"
            }, {
                data: "Icon"
            }, {
                data: "Area"
            }, {
                data: "Controller"
            }, {
                data: "URL"
            }],
            select: true,
            buttons: []
        });
    }
})(this)