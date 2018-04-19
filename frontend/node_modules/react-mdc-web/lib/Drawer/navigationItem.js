'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

var _detail = require('../List/detail');

var _detail2 = _interopRequireDefault(_detail);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var navigationItem = function navigationItem(component, isTemporary) {
  var _component$props = component.props,
      className = _component$props.className,
      selected = _component$props.selected,
      children = _component$props.children;

  var classes = (0, _classnames2.default)('mdc-list-item', {
    'mdc-permanent-drawer--selected': selected && !isTemporary,
    'mdc-temporary-drawer--selected': selected && isTemporary
  }, className);

  var childs = _react2.default.Children.map(children, function (child) {
    return (0, _detail2.default)(child);
  });

  return _react2.default.cloneElement(component, { className: classes, children: childs });
};
exports.default = navigationItem;