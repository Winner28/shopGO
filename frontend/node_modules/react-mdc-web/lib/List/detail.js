'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

var _Icon = require('../Icon');

var _Icon2 = _interopRequireDefault(_Icon);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var detail = function detail(component) {
  var start = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;

  var modificatorClass = start ? 'mdc-list-item__start-detail' : 'mdc-list-item__end-detail';
  if (component.type === _Icon2.default) {
    var className = component.props.className;

    var classes = (0, _classnames2.default)(modificatorClass, className);
    return _react2.default.cloneElement(component, { className: classes, 'aria-hidden': true });
  }
  return component;
};
exports.default = detail;