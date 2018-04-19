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

var _ListDivider = require('./ListDivider');

var _ListDivider2 = _interopRequireDefault(_ListDivider);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  dense: _propTypes2.default.bool
};

var List = function List(_ref) {
  var className = _ref.className,
      children = _ref.children,
      dense = _ref.dense,
      otherProps = _objectWithoutProperties(_ref, ['className', 'children', 'dense']);

  var childs = _react.Children.map(children, function (child) {
    if (child.type === _ListDivider2.default) {
      return (0, _react.cloneElement)(child, { isListItem: true });
    }
    return child;
  });
  return _react2.default.createElement(
    'ul',
    _extends({
      className: (0, _classnames2.default)('mdc-list', {
        'mdc-list--dense': dense
      }, className)
    }, otherProps),
    childs
  );
};
List.propTypes = propTypes;
exports.default = List;