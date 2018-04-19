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

var _Icon = require('../Icon');

var _Icon2 = _interopRequireDefault(_Icon);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var ROOT = 'mdc-fab';

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node,
  mini: _propTypes2.default.bool
};

var Fab = function Fab(_ref) {
  var className = _ref.className,
      mini = _ref.mini,
      otherProps = _objectWithoutProperties(_ref, ['className', 'mini']);

  var classes = (0, _classnames3.default)(ROOT, _defineProperty({}, ROOT + '--mini', mini), className);

  var children = _react.Children.map(otherProps.children, function (child) {
    if (child.type === _Icon2.default) {
      var childClasses = child.props.className;
      return (0, _react.cloneElement)(child, {
        className: (0, _classnames3.default)(childClasses, 'mdc-fab__icon')
      });
    }
    return child;
  });

  return _react2.default.createElement(
    'button',
    _extends({
      className: classes
    }, otherProps),
    children
  );
};
Fab.propTypes = propTypes;
exports.default = Fab;