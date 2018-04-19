'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var ROOT = 'mdc-form-field';
var ALIGN_END = ROOT + '--align-end';

var propTypes = {
  alignEnd: _propTypes2.default.bool,
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  id: _propTypes2.default.string.isRequired
};

var FormField = function FormField(_ref) {
  var alignEnd = _ref.alignEnd,
      className = _ref.className,
      children = _ref.children,
      id = _ref.id,
      otherProps = _objectWithoutProperties(_ref, ['alignEnd', 'className', 'children', 'id']);

  var childs = _react.Children.map(children, function (child) {
    if (child.type === 'label') {
      return (0, _react.cloneElement)(child, { htmlFor: id });
    }
    return (0, _react.cloneElement)(child, { id: id });
  });

  return _react2.default.createElement(
    'div',
    _extends({
      className: (0, _classnames3.default)(ROOT, _defineProperty({}, ALIGN_END, alignEnd), className)
    }, otherProps),
    childs
  );
};
FormField.propTypes = propTypes;
exports.default = FormField;