'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

var _detail = require('../List/detail');

var _detail2 = _interopRequireDefault(_detail);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node
};

var ListItem = function ListItem(_ref) {
  var className = _ref.className,
      children = _ref.children,
      otherProps = _objectWithoutProperties(_ref, ['className', 'children']);

  var childs = _react.Children.map(children, function (child, index) {
    return (0, _detail2.default)(child, index === 0);
  });
  return _react2.default.createElement(
    'li',
    _extends({
      className: (0, _classnames2.default)('mdc-list-item', className)
    }, otherProps),
    childs
  );
};
ListItem.propTypes = propTypes;
exports.default = ListItem;