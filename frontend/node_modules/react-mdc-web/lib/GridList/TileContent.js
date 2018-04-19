'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  component: _propTypes2.default.string
};

var defaultProps = {
  component: 'img'
};

var TileContent = function TileContent(_ref) {
  var className = _ref.className,
      component = _ref.component,
      otherProps = _objectWithoutProperties(_ref, ['className', 'component']);

  var classes = (0, _classnames2.default)('mdc-grid-tile__primary-content', className);
  return (0, _react.createElement)(component, _extends({
    className: classes
  }, otherProps));
};

TileContent.propTypes = propTypes;
TileContent.defaultProps = defaultProps;
exports.default = TileContent;