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

var _Icon = require('../Icon/Icon');

var _Icon2 = _interopRequireDefault(_Icon);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  children: _propTypes2.default.node,
  className: _propTypes2.default.string
};

var TileSecondary = function TileSecondary(_ref) {
  var className = _ref.className,
      children = _ref.children,
      otherProps = _objectWithoutProperties(_ref, ['className', 'children']);

  var childs = _react.Children.map(children, function (child) {
    if (child.type === _Icon2.default) {
      var childClassName = child.props.className;

      var classes = (0, _classnames2.default)('mdc-grid-tile__icon', childClassName);
      return _react2.default.cloneElement(child, { className: classes });
    }
    return child;
  });
  return _react2.default.createElement(
    'span',
    _extends({
      className: (0, _classnames2.default)('mdc-grid-tile__secondary', className)
    }, otherProps),
    childs
  );
};
TileSecondary.propTypes = propTypes;
exports.default = TileSecondary;